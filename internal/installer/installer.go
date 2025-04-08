package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/registry"
	spinner "github.com/moq77111113/kite/internal/vendors"
)

func InstallAllModules(modules []*registry.Module, destDir string) error {
	if len(modules) == 0 {
		return nil
	}

	s := spinner.StartWithMessage("Installing modules")
	defer s.Stop()

	for i, mod := range modules {
		s.UpdateMessagef("Installing module %d/%d: %s", i+1, len(modules), mod.Name)
		if err := installModule(mod, destDir); err != nil {
			return fmt.Errorf("could not install module %s: %w", mod.Name, err)
		}
	}

	return nil
}

func installModule(module *registry.Module, destDir string) error {
	moduleDestDir := filepath.Join(destDir, module.Name)
	if err := os.MkdirAll(moduleDestDir, 0755); err != nil {
		return fmt.Errorf("could not create module directory: %w", err)
	}

	for _, file := range module.Files {
		
		filePath := filepath.Join(moduleDestDir, file.Path)
		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return fmt.Errorf("could not create directory for file %s: %w", filePath, err)
		}

		if err := writeFile(filePath, file.Content); err != nil {
			return fmt.Errorf("could not write file %s: %w", filePath, err)
		}
	}

	return nil
}


func writeFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}