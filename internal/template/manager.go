package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/config"
	"github.com/moq77111113/kite/internal/registry"
)

// Manager handles template operations
type Manager struct {
	config    *config.Config
	client    registry.Client
	installer Installer
}

func NewManager(cfg *config.Config, client registry.Client) *Manager {
	return &Manager{
		config:    cfg,
		client:    client,
		installer: NewInstaller(),
	}
}

// Add downloads and installs a template
func (m *Manager) Add(name string) error {
	// Check if already installed
	if _, exists := m.config.GetTemplate(name); exists {
		return fmt.Errorf("template %s is already installed", name)
	}

	// Fetch from registry
	template, err := m.client.GetTemplate(name)
	if err != nil {
		return fmt.Errorf("failed to fetch template: %w", err)
	}

	// Install
	destPath := filepath.Join(m.config.Path, name)
	if err := m.installer.Install(template, destPath); err != nil {
		return fmt.Errorf("failed to install template: %w", err)
	}

	// Update config
	m.config.AddTemplate(name, template.Version)
	if err := config.Save(m.config); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// Remove uninstalls a template
func (m *Manager) Remove(name string) error {
	// Check if installed
	if _, exists := m.config.GetTemplate(name); !exists {
		return fmt.Errorf("template %s is not installed", name)
	}

	// Remove directory
	destPath := filepath.Join(m.config.Path, name)
	if err := os.RemoveAll(destPath); err != nil {
		return fmt.Errorf("failed to remove directory: %w", err)
	}

	// Update config
	m.config.RemoveTemplate(name)
	if err := config.Save(m.config); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// CheckUpdate checks if an update is available for a template
func (m *Manager) CheckUpdate(name string) (*UpdateInfo, error) {
	// Get installed version
	installed, exists := m.config.GetTemplate(name)
	if !exists {
		return nil, fmt.Errorf("template %s is not installed", name)
	}

	// Fetch latest from registry
	latest, err := m.client.GetTemplate(name)
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
