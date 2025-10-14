package initcmd

import (
	"fmt"

	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

// Options contains initialization options
type Options struct {
	Registry string
	Path     string
	Force    bool
}

// Run executes the initialization
func Run(opts Options) error {
	existing, err := loadExisting(opts.Force)
	if err != nil {
		return err
	}

	registry, path, err := getValues(opts, existing)
	if err != nil {
		return err
	}

	return saveConfig(existing, registry, path)
}

func getValues(opts Options, existing *config.Config) (string, string, error) {
	if opts.Registry != "" || opts.Path != "" {
		return fromOptions(opts, existing)
	}
	return interactive(existing)
}

func interactive(existing *config.Config) (string, string, error) {
	showWelcome()

	registryType, err := AskRegistryType()
	if err != nil {
		return "", "", err
	}

	registry, err := AskRegistryURL(registryType)
	if err != nil {
		return "", "", err
	}

	defaultPath := config.DefaultPath
	if existing != nil {
		defaultPath = existing.Path
	}

	path, err := AskPath(defaultPath)
	if err != nil {
		return "", "", err
	}

	return registry, path, nil
}

func fromOptions(opts Options, existing *config.Config) (string, string, error) {
	registry := opts.Registry
	path := opts.Path

	if registry == "" && existing != nil {
		registry = existing.Registry
	}
	if registry == "" {
		return "", "", fmt.Errorf("--registry flag is required")
	}

	if path == "" && existing != nil {
		path = existing.Path
	}
	if path == "" {
		path = config.DefaultPath
	}

	return registry, path, nil
}
