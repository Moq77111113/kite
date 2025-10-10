package container

import (
	"github.com/moq77111113/kite/internal/config"
	"github.com/moq77111113/kite/internal/registry"
	"github.com/moq77111113/kite/internal/template"
)

// Container holds all application dependencies
type Container struct {
	config    *config.Config
	client    registry.Client
	installer template.Installer
	manager   *template.Manager
}

// New creates a new container with all dependencies
func New() (*Container, error) {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return NewWithConfig(cfg), nil
}

// NewWithConfig creates a container with a provided config
func NewWithConfig(cfg *config.Config) *Container {
	// Create registry client based on config
	var client registry.Client
	// TODO: Check if registry is URL or file path
	// For now, use mock client
	client = registry.NewMockClient()

	// Create installer
	installer := template.NewInstaller()

	// Create template manager
	manager := template.NewManager(cfg, client)

	return &Container{
		config:    cfg,
		client:    client,
		installer: installer,
		manager:   manager,
	}
}

// Config returns the configuration
func (c *Container) Config() *config.Config {
	return c.config
}

// Client returns the registry client
func (c *Container) Client() registry.Client {
	return c.client
}

// Installer returns the template installer
func (c *Container) Installer() template.Installer {
	return c.installer
}

// Manager returns the template manager
func (c *Container) Manager() *template.Manager {
	return c.manager
}
