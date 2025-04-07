package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/registry"
)

func InstallModule(module *registry.RegistryItem, destDir string) error {
	moduleDestDir := filepath.Join(destDir, module.Name)
	if err := os.MkdirAll(moduleDestDir, 0755); err != nil {
		return fmt.Errorf("could not create module directory: %w", err)
	}

	for _, file := range module.Files {
		 
		sourcePath := filepath.Join("./__registry__", file)

		if err := copyFile(sourcePath, filepath.Join(moduleDestDir, file)); err != nil {
			return fmt.Errorf("could not copy file: %w", err)
		}
	}

	return nil
}
func copyFile(sourcePath, destPath string) error {
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("could not read source file: %w", err)
	}

	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	if err := os.WriteFile(destPath, data, 0644); err != nil {
		return fmt.Errorf("could not write to destination file: %w", err)
	}

	return nil
}
