package initcmd

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/config"
)

func saveConfig(existing *config.Config, registry, path string) error {
	if existing != nil {
		return updateConfig(existing, registry, path)
	}
	return createConfig(registry, path)
}

func updateConfig(cfg *config.Config, registry, path string) error {
	cfg.Registry = registry
	cfg.Path = path
	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}
	showUpdateSuccess(cfg, registry, path)
	return nil
}

func createConfig(registry, path string) error {
	if _, err := config.Init(registry, path); err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}
	showCreateSuccess(registry, path)
	return nil
}

func loadExisting(force bool) (*config.Config, error) {
	if !config.Exists() {
		return nil, nil
	}

	existing, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load existing config: %w", err)
	}

	if !force {
		shouldUpdate, err := promptUpdate(existing)
		if err != nil {
			return nil, err
		}
		if !shouldUpdate {
			return nil, fmt.Errorf("cancelled by user")
		}
	}

	return existing, nil
}
