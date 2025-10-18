package initcmd

import (
	"fmt"

	initapp "github.com/moq77111113/kite/internal/application/init"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

type Options struct {
	Registry string
	Path     string
	Force    bool
}

func Run(opts Options) error {
	if !opts.Force && config.Exists("") {
		existing, err := config.Load("")
		if err != nil {
			return fmt.Errorf("failed to load existing config: %w", err)
		}

		shouldUpdate, err := promptUpdate(existing)
		if err != nil {
			return err
		}
		if !shouldUpdate {
			return fmt.Errorf("cancelled by user")
		}
	}

	registry, path, err := getValues(opts)
	if err != nil {
		return err
	}

	initSvc := initapp.New()

	result, err := initSvc.Execute(initapp.Request{
		Registry: registry,
		Path:     path,
		Force:    opts.Force,
	})
	if err != nil {
		return err
	}

	if result.Created {
		showCreateSuccess(result.Registry, result.Path)
	} else {
		cfg, _ := config.Load("")
		showUpdateSuccess(cfg, result.Registry, result.Path)
	}

	return nil
}

func getValues(opts Options) (string, string, error) {
	if opts.Registry != "" && opts.Path != "" {
		return opts.Registry, opts.Path, nil
	}

	if opts.Registry != "" || opts.Path != "" {
		return fromOptions(opts)
	}

	return interactive()
}

func interactive() (string, string, error) {
	showWelcome()

	registryType, err := AskRegistryType()
	if err != nil {
		return "", "", err
	}

	registry, err := AskRegistryURL(registryType)
	if err != nil {
		return "", "", err
	}

	existing, _ := config.Load("")
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

func fromOptions(opts Options) (string, string, error) {
	registry := opts.Registry
	path := opts.Path

	if registry == "" {
		return "", "", fmt.Errorf("--registry flag is required")
	}

	if path == "" {
		existing, _ := config.Load("")
		if existing != nil {
			path = existing.Path
		} else {
			path = config.DefaultPath
		}
	}

	return registry, path, nil
}
