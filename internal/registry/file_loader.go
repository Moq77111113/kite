package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type FileLoader struct{}

func (f *FileLoader) Load(path string) (*Registry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var r Registry
	switch filepath.Ext(path) {
	case ".json":
		err = json.Unmarshal(data, &r)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(data, &r)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", path)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid registry format: %w", err)
	}

	return &r, nil
}