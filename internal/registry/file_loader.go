package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/config"
	spinner "github.com/moq77111113/kite/internal/vendors"
	"gopkg.in/yaml.v3"
)

type FileLoader struct{}

func (h *FileLoader) LoadIndex(config config.Config) (*Registry, error) {
	var result *Registry
	err := spinner.WithContext(fmt.Sprintf("Loading registry index from %s", config.Registry), func() error {

	data, err := os.ReadFile(config.Registry)
	if err != nil {
		return  err
	}

	var r Registry
	switch filepath.Ext(config.Registry) {
	case ".json":
		err = json.Unmarshal(data, &r)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &r)
	default:
		return fmt.Errorf("unsupported file type: %s", config.Registry)
	}

	if err != nil {
		return  fmt.Errorf("invalid registry format: %w", err)
	}

	result = &r
	return nil
	})
	return result, err
}


func (h *FileLoader) LoadModules(config config.Config, names []string) ([]*Module, error) {
	modules := make([]*Module, 0, len(names))
	if len(names) == 0 {
		return modules, nil
	}

	s := spinner.StartWithMessage("Loading modules")
	defer s.Stop()

	parentDir := filepath.Dir(config.Registry)
	if _, err := os.Stat(filepath.Join(parentDir, "__registry__")); err == nil {
		config.Registry = filepath.Join(parentDir, "__registry__")
	}

	flavorPath := filepath.Join(config.Registry, config.Flavor)

	for i, name := range names {
		s.UpdateMessagef("Loading module %s (%d/%d)", name, i+1, len(names))
		u := filepath.Join(flavorPath, "modules", name+".json")
		data, err := os.ReadFile(u)
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}

		var m Module
		if err = json.Unmarshal(data, &m); err != nil {
			return nil, fmt.Errorf("invalid module format: %w", err)
		}
		modules = append(modules, &m)
	}

	return modules, nil
}