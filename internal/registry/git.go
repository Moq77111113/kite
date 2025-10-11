package registry

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/git"
	"gopkg.in/yaml.v3"
)

// GitClient implements registry.Client using a Git repository
type GitClient struct {
	repoURL   string
	cachePath string
	git       git.Client
}

// NewGitClient creates a new Git-based registry client
func NewGitClient(repoURL string, git git.Client) (Client, error) {
	var cachePath string
	var isLocalRepo bool

	// Check if repoURL is already a local Git repository
	if isLocalPath(repoURL) && git.IsCloned(repoURL) {
		// Use the local path directly
		cachePath = repoURL
		isLocalRepo = true
	} else {
		// Remote repo - will be cloned to cache
		cachePath = getCachePath(repoURL)
		isLocalRepo = false
	}

	client := &GitClient{
		repoURL:   repoURL,
		cachePath: cachePath,
		git:       git,
	}

	// Clone or pull repository (skip if local)
	if !isLocalRepo {
		if err := client.ensureRepository(); err != nil {
			return nil, fmt.Errorf("failed to initialize git registry: %w", err)
		}
	}

	return client, nil
}

// ListTemplates returns all templates from the Git registry
func (c *GitClient) ListTemplates() ([]TemplateSummary, error) {
	templatesDir := filepath.Join(c.cachePath, "templates")

	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read templates directory: %w", err)
	}

	var templates []TemplateSummary
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		templatePath := filepath.Join(templatesDir, entry.Name())
		metadata, err := c.readKiteYAML(templatePath)
		if err != nil {
			// Skip templates with invalid kite.yaml
			continue
		}

		templates = append(templates, TemplateSummary{
			Name:        metadata.Name,
			Description: metadata.Description,
			Version:     metadata.Version,
			Tags:        metadata.Tags,
			Author:      metadata.Author,
		})
	}

	return templates, nil
}

// GetTemplate returns a specific template by name
func (c *GitClient) GetTemplate(name string) (*TemplateDetailResponse, error) {
	templatePath := filepath.Join(c.cachePath, "templates", name)

	// Check if template exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("template %s not found", name)
	}

	// Read metadata
	metadata, err := c.readKiteYAML(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template metadata: %w", err)
	}

	// Read all files
	files, err := c.readTemplateFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template files: %w", err)
	}

	// Read README if exists
	readme := ""
	readmePath := filepath.Join(templatePath, "README.md")
	if data, err := os.ReadFile(readmePath); err == nil {
		readme = string(data)
	}

	return &TemplateDetailResponse{
		Name:        metadata.Name,
		Version:     metadata.Version,
		Description: metadata.Description,
		Files:       files,
		Variables:   metadata.Variables,
		Readme:      readme,
	}, nil
}

// ensureRepository clones or updates the Git repository
func (c *GitClient) ensureRepository() error {
	if c.git.IsCloned(c.cachePath) {
		// Repository exists, pull updates
		return c.git.Pull(c.cachePath)
	}

	// Repository doesn't exist, clone it
	if err := os.MkdirAll(filepath.Dir(c.cachePath), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	return c.git.Clone(c.repoURL, c.cachePath)
}

// readKiteYAML reads and parses kite.yaml from a template directory
func (c *GitClient) readKiteYAML(templatePath string) (*KiteYAML, error) {
	kiteYAMLPath := filepath.Join(templatePath, "kite.yaml")

	data, err := os.ReadFile(kiteYAMLPath)
	if err != nil {
		return nil, fmt.Errorf("kite.yaml not found: %w", err)
	}

	var metadata KiteYAML
	if err := yaml.Unmarshal(data, &metadata); err != nil {
		return nil, fmt.Errorf("invalid kite.yaml: %w", err)
	}

	// Validate required fields
	if metadata.Name == "" || metadata.Version == "" {
		return nil, fmt.Errorf("kite.yaml missing required fields (name, version)")
	}

	return &metadata, nil
}

// readTemplateFiles reads all files from a template directory
func (c *GitClient) readTemplateFiles(templatePath string) ([]TemplateFile, error) {
	var files []TemplateFile

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories, kite.yaml, and README.md
		if info.IsDir() || info.Name() == "kite.yaml" || info.Name() == "README.md" {
			return nil
		}

		// Skip .git directory but allow other dotfiles (like .gitlab-ci.yml)
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		// Get relative path from template directory
		relPath, err := filepath.Rel(templatePath, path)
		if err != nil {
			return err
		}

		files = append(files, TemplateFile{
			Path:    relPath,
			Content: string(content),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// getCachePath returns the cache path for a Git repository
func getCachePath(repoURL string) string {
	// Hash the repo URL to create a unique cache directory
	hash := sha256.Sum256([]byte(repoURL))
	hashStr := fmt.Sprintf("%x", hash[:8])

	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return filepath.Join(homeDir, ".kite", "cache", "registry", hashStr)
}

// KiteYAML represents the kite.yaml metadata file
type KiteYAML struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Version     string     `yaml:"version"`
	Author      string     `yaml:"author"`
	Tags        []string   `yaml:"tags"`
	Variables   []Variable `yaml:"variables,omitempty"`
}

// isLocalPath checks if a path is a local filesystem path
func isLocalPath(path string) bool {
	return filepath.IsAbs(path) ||
		   len(path) > 0 && path[0] == '.' ||
		   (!hasPrefix(path, "http://") &&
		    !hasPrefix(path, "https://") &&
		    !hasPrefix(path, "git@"))
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
