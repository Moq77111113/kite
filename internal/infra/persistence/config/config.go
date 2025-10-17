package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Config represents the kite.yaml configuration file
type Config struct {
	Version   string              `yaml:"version"`
	Registry  string              `yaml:"registry"`
	Path      string              `yaml:"path"`
	Templates map[string]Template `yaml:"templates"`
}

// Template represents an installed template
type Template struct {
	Version   string `yaml:"version"`
	Installed int64  `yaml:"installed"`
}

const (
	DefaultPath        = "./"
	ConfigFileName     = "kite.yaml"
	GlobalConfigName   = "config.yaml"
)

// GetConfigPath returns the path to the kite config file
// Priority:
//   1. --config flag (passed as customPath)
//   2. ./kite.json in current directory (project-specific)
//   3. ~/.kite/config.json (global fallback)
func GetConfigPath(customPath string) string {
	// 1. Custom path from flag has highest priority
	if customPath != "" {
		return customPath
	}

	// 2. Check for kite.json in current directory (project-specific)
	projectConfig := filepath.Join(".", ConfigFileName)
	if _, err := os.Stat(projectConfig); err == nil {
		return projectConfig
	}

	// 3. Fallback to global config in ~/.kite/config.json
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// If can't get home dir, use current directory as last resort
		return projectConfig
	}

	return filepath.Join(homeDir, ".kite", GlobalConfigName)
}

// GetConfigPathForLoad is like GetConfigPath but doesn't create anything
func GetConfigPathForLoad(customPath string) (string, error) {
	path := GetConfigPath(customPath)

	// Verify file exists
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("config not found at %s (run 'kite init' first)", path)
	}

	return path, nil
}

// Load reads the kite configuration file
// customPath: optional --config flag value
func Load(customPath string) (*Config, error) {
	configPath := GetConfigPath(customPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config not found at %s (run 'kite init' first)", configPath)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("invalid config at %s: %w", configPath, err)
	}

	return &config, nil
}

// Save writes the configuration to the config file
// customPath: optional --config flag value
func Save(config *Config, customPath string) error {
	configPath := GetConfigPath(customPath)

	// Ensure directory exists
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// Exists checks if kite config file exists
// customPath: optional --config flag value
func Exists(customPath string) bool {
	configPath := GetConfigPath(customPath)
	_, err := os.Stat(configPath)
	return err == nil
}

// Init creates a new kite configuration file
// Creates in current directory (./kite.yaml) if --local flag, otherwise in ~/.kite/config.yaml
func Init(registry, path string, local bool, customPath string) (*Config, error) {
	if path == "" {
		path = DefaultPath
	}
	if registry == "" {
		return nil, fmt.Errorf("registry cannot be empty")
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

	var configLocation string
	if local {

		configLocation = filepath.Join(".", ConfigFileName)
	} else {
		configLocation = customPath
	}

	// Save the config
	if err := Save(config, configLocation); err != nil {
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
