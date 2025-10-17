package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

// Repository handles template operations
type Repository struct {
	config    *config.Config
	client    registry.Client
	installer Installer
}

func NewRepository(cfg *config.Config, client registry.Client) *Repository {
	return &Repository{
		config:    cfg,
		client:    client,
		installer: NewInstaller(),
	}
}

// CheckConflict checks if a template directory already exists
func (r *Repository) CheckConflict(destPath string) (bool, error) {
	_, err := os.Stat(destPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Add downloads and installs a template
func (r *Repository) Add(name, customPath string) error {
	template, err := r.client.GetTemplate(name)
	if err != nil {
		return fmt.Errorf("failed to fetch template: %w", err)
	}

	destPath := filepath.Join(r.config.Path, name)
	if customPath != "" {
		destPath = customPath
	}

	if err := r.installer.Install(template, destPath); err != nil {
		return fmt.Errorf("failed to install template: %w", err)
	}

	r.config.AddTemplate(name, template.Version)
	if err := config.Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// Remove uninstalls a template
func (r *Repository) Remove(name string) error {
	if _, exists := r.config.GetTemplate(name); !exists {
		return fmt.Errorf("template %s is not installed", name)
	}

	destPath := filepath.Join(r.config.Path, name)
	if err := os.RemoveAll(destPath); err != nil {
		return fmt.Errorf("failed to remove directory: %w", err)
	}

	r.config.RemoveTemplate(name)
	if err := config.Save(r.config, ""); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// CheckUpdate checks if an update is available for a template
func (r *Repository) CheckUpdate(name string) (*UpdateInfo, error) {
	installed, exists := r.config.GetTemplate(name)
	if !exists {
		return nil, fmt.Errorf("template %s is not installed", name)
	}

	latest, err := r.client.GetTemplate(name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch template: %w", err)
	}

	return &UpdateInfo{
		Name:           name,
		CurrentVersion: installed.Version,
		LatestVersion:  latest.Version,
		UpdateAvailable: installed.Version != latest.Version,
	}, nil
}

type UpdateInfo struct {
	Name            string
	CurrentVersion  string
	LatestVersion   string
	UpdateAvailable bool
}
