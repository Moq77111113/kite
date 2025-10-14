package git

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/registry"
	"gopkg.in/yaml.v3"
)

// KiteYAML represents the kite.yaml metadata file
type KiteYAML struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Version     string              `yaml:"version"`
	Author      string              `yaml:"author"`
	Tags        []string            `yaml:"tags"`
	Variables   []registry.Variable `yaml:"variables,omitempty"`
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
func (c *GitClient) readTemplateFiles(templatePath string) ([]registry.TemplateFile, error) {
	var files []registry.TemplateFile

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

		files = append(files, registry.TemplateFile{
			Path:    relPath,
			Content: string(content),
		})

		return nil
	})

	return files, err
}
