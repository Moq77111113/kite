package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/config"
	"gopkg.in/yaml.v3"
)

type FileLoader struct{}

func (h *FileLoader) LoadIndex(config config.Config) (*Registry, error) {
	data, err := os.ReadFile(config.Registry)
	if err != nil {
		return nil, err
	}

	var r Registry
	switch filepath.Ext(config.Registry) {
	case ".json":
		err = json.Unmarshal(data, &r)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &r)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", config.Registry)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid registry format: %w", err)
	}

	return &r, nil
}


func (h *FileLoader) LoadModules(config config.Config, names []string) ([]*Module, error) {
	modules := make([]*Module, 0, len(names))
	if len(names) == 0 {
		return modules, nil
	}

	// look for a "__registry__" directory in the config.registry parent dir
	parentDir := filepath.Dir(config.Registry)
	if _, err := os.Stat(filepath.Join(parentDir, "__registry__")); err == nil {
		config.Registry = filepath.Join(parentDir, "__registry__")
	}

	flavorPath := filepath.Join(config.Registry, config.Flavor)

	for _, name := range names {
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