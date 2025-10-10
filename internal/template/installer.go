package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/registry"
)

// Installer handles installing templates to the filesystem
type Installer interface {
	Install(template *registry.TemplateDetailResponse, destPath string) error
}

type installer struct{}

func NewInstaller() Installer {
	return &installer{}
}

func (i *installer) Install(template *registry.TemplateDetailResponse, destPath string) error {
	// Create template directory
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Write all files
	for _, file := range template.Files {
		if err := i.writeFile(destPath, file); err != nil {
			return err
		}
	}

	return nil
}

func (i *installer) writeFile(basePath string, file registry.TemplateFile) error {
	filePath := filepath.Join(basePath, file.Path)

	// Create parent directory if needed
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write file
	if err := os.WriteFile(filePath, []byte(file.Content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}
