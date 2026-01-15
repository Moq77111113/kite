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

func (r *Registry) Add(id, version string) error {
	r.config.AddKit(id, version)

	if err := Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}

func (r *Registry) Remove(id string) error {
	r.config.RemoveKit(id)

	if err := Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}

func (r *Registry) Get(id string) (*models.InstalledKit, error) {
	kitInfo, exists := r.config.GetKit(id)
	if !exists {
		return nil, fmt.Errorf("kit %s not found", id)
	}

	return &models.InstalledKit{
		ID:        id,
		Version:   kitInfo.Version,
		Installed: time.Unix(kitInfo.Installed, 0),
	}, nil
}

func (r *Registry) List() []models.InstalledKit {
	var installed []models.InstalledKit
	for id, kitInfo := range r.config.Kits {
		installed = append(installed, models.InstalledKit{
			ID:        id,
			Version:   kitInfo.Version,
			Installed: time.Unix(kitInfo.Installed, 0),
		})
	}
	return installed
}
