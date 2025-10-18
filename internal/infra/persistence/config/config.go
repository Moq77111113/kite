package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version  string         `yaml:"version"`
	Registry string         `yaml:"registry"`
	Path     string         `yaml:"path"`
	Kits     map[string]Kit `yaml:"kits"`
}

type Kit struct {
	Version   string `yaml:"version"`
	Installed int64  `yaml:"installed"`
}

const (
	DefaultPath      = "./"
	ConfigFileName   = "kite.yaml"
	GlobalConfigName = "config.yaml"
)

func GetConfigPath(customPath string) string {
	if customPath != "" {
		return customPath
	}

	projectConfig := filepath.Join(".", ConfigFileName)
	if _, err := os.Stat(projectConfig); err == nil {
		return projectConfig
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return projectConfig
	}

	return filepath.Join(homeDir, ".kite", GlobalConfigName)
}

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

func Save(config *Config, customPath string) error {
	configPath := GetConfigPath(customPath)

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

func Exists(customPath string) bool {
	configPath := GetConfigPath(customPath)
	_, err := os.Stat(configPath)
	return err == nil
}

func Init(registry, path string, local bool, customPath string) (*Config, error) {
	if path == "" {
		path = DefaultPath
	}
	if registry == "" {
		return nil, fmt.Errorf("registry cannot be empty")
	}

	config := &Config{
		Version:  "1.0.0",
		Registry: registry,
		Path:     path,
		Kits:     make(map[string]Kit),
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

func (c *Config) AddKit(name, version string) {
	if c.Kits == nil {
		c.Kits = make(map[string]Kit)
	}
	c.Kits[name] = Kit{
		Version:   version,
		Installed: time.Now().Unix(),
	}
}

func (c *Config) RemoveKit(name string) {
	delete(c.Kits, name)
}

func (c *Config) GetKit(name string) (Kit, bool) {
	t, ok := c.Kits[name]
	return t, ok
}
