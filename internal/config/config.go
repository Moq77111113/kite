package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Config represents the kite.json configuration file
type Config struct {
	Version   string              `json:"version"`
	Registry  string              `json:"registry"`
	Path      string              `json:"path"`
	Templates map[string]Template `json:"templates"`
}

// Template represents an installed template
type Template struct {
	Version   string `json:"version"`
	Installed int64  `json:"installed"`
}

const (
	DefaultRegistry = "https://api.kite.sh"
	DefaultPath     = "./infrastructure"
	ConfigFileName  = "kite.json"
)

// Load reads the kite.json configuration file from the current directory
func Load() (*Config, error) {
	configPath := filepath.Join(".", ConfigFileName)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Save writes the configuration to kite.json in the current directory
func Save(config *Config) error {
	configPath := filepath.Join(".", ConfigFileName)

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// Exists checks if kite.json exists in the current directory
func Exists() bool {
	configPath := filepath.Join(".", ConfigFileName)
	_, err := os.Stat(configPath)
	return err == nil
}

// Init creates a new kite.json configuration file
func Init(registry, path string) (*Config, error) {
	if registry == "" {
		registry = DefaultRegistry
	}
	if path == "" {
		path = DefaultPath
	}

	config := &Config{

		Version:   "1.0.0",
		Registry:  registry,
		Path:      path,
		Templates: make(map[string]Template),
	}


	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	// Save the config
	if err := Save(config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) AddTemplate(name, version string) {
	if c.Templates == nil {
		c.Templates = make(map[string]Template)
	}
	c.Templates[name] = Template{
		Version:   version,
		Installed: time.Now().Unix(),
	}
}

func (c *Config) RemoveTemplate(name string) {
	delete(c.Templates, name)
}


func (c *Config) GetTemplate(name string) (Template, bool) {
	t, ok := c.Templates[name]
	return t, ok
}
