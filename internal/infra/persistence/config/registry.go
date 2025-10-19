package config

import (
	"fmt"
	"time"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Registry struct {
	config *Config
}

func NewKitRegistry(cfg *Config) models.KitRegistry {
	return &Registry{
		config: cfg,
	}
}

func (r *Registry) Add(name, version string) error {
	r.config.AddKit(name, version)

	if err := Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}

func (r *Registry) Remove(name string) error {
	r.config.RemoveKit(name)

	if err := Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}

func (r *Registry) Get(name string) (*models.InstalledKit, error) {
	kitInfo, exists := r.config.GetKit(name)
	if !exists {
		return nil, fmt.Errorf("kit %s not found", name)
	}

	return &models.InstalledKit{
		Name:      name,
		Version:   kitInfo.Version,
		Installed: time.Unix(kitInfo.Installed, 0),
	}, nil
}

func (r *Registry) List() []models.InstalledKit {
	var installed []models.InstalledKit
	for name, kitInfo := range r.config.Kits {
		installed = append(installed, models.InstalledKit{
			Name:      name,
			Version:   kitInfo.Version,
			Installed: time.Unix(kitInfo.Installed, 0),
		})
	}
	return installed
}
