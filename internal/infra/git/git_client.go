package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/api/registry/v1"
)

// GitClient implements registry.Client using a Git repository
type GitClient struct {
	repoURL   string
	cachePath string
	git       Client // Local git wrapper client
}

// NewGitClient creates a new Git-based registry client
func NewGitClient(repoURL string, gitClient Client) (registry.Client, error) {
	var cachePath string
	var isLocalRepo bool

	// Check if repoURL is already a local Git repository
	if isLocalPath(repoURL) && gitClient.IsCloned(repoURL) {
		cachePath = repoURL
		isLocalRepo = true
	} else {
		cachePath = getCachePath(repoURL)
		isLocalRepo = false
	}

	client := &GitClient{
		repoURL:   repoURL,
		cachePath: cachePath,
		git:       gitClient,
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
func (c *GitClient) ListTemplates() ([]registry.TemplateSummary, error) {
	entries, err := os.ReadDir(c.cachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read registry directory: %w", err)
	}

	var templates []registry.TemplateSummary
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skipDirs := map[string]bool{".git": true, ".github": true, "docs": true, "examples": true}
		if skipDirs[entry.Name()] {
			continue
		}

		templatePath := filepath.Join(c.cachePath, entry.Name())
		metadata, err := c.readKiteYAML(templatePath)
		if err != nil {
			continue // Skip directories without valid kite.yaml
		}

		templates = append(templates, registry.TemplateSummary{
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
func (c *GitClient) GetTemplate(name string) (*registry.TemplateDetailResponse, error) {
	templatePath := filepath.Join(c.cachePath, name)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("template %s not found", name)
	}

	metadata, err := c.readKiteYAML(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template metadata: %w", err)
	}

	files, err := c.readTemplateFiles(templatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template files: %w", err)
	}

	readme := ""
	readmePath := filepath.Join(templatePath, "README.md")
	if data, err := os.ReadFile(readmePath); err == nil {
		readme = string(data)
	}

	return &registry.TemplateDetailResponse{
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
		return c.git.Pull(c.cachePath)
	}

	if err := os.MkdirAll(filepath.Dir(c.cachePath), 0755); err != nil {
		return fmt.Errorf("failed to create cache directory: %w", err)
	}

	return c.git.Clone(c.repoURL, c.cachePath)
}
