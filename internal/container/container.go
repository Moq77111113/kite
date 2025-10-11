package container

import (
	"fmt"

	"github.com/moq77111113/kite/internal/config"
	"github.com/moq77111113/kite/internal/git"
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
	client := createRegistryClient(cfg.Registry)

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

// createRegistryClient creates the appropriate registry client based on registry URL
func createRegistryClient(registryURL string) registry.Client {
	registryType := registry.DetectRegistryType(registryURL)

	switch registryType {
	case registry.RegistryTypeGit:
		gitClient := git.NewClient()
		client, err := registry.NewGitClient(registryURL, gitClient)
		if err != nil {
			// Fall back to mock client on error
			fmt.Printf("Warning: Failed to initialize Git registry, using mock: %v\n", err)
			return registry.NewMockClient()
		}
		return client

	case registry.RegistryTypeHTTP:
		return registry.NewHTTPClient(registryURL)

	case registry.RegistryTypeLocal:
		// Check if it's a Git repository
		gitClient := git.NewClient()
		if gitClient.IsCloned(registryURL) {
			// It's a local Git repo, use GitClient
			client, err := registry.NewGitClient(registryURL, gitClient)
			if err != nil {
				fmt.Printf("Warning: Failed to initialize local Git registry, using mock: %v\n", err)
				return registry.NewMockClient()
			}
			return client
		}
		// TODO: Implement file-based LocalClient
		fmt.Println("Warning: File-based local registry not yet implemented, using mock")
		return registry.NewMockClient()

	default:
		// Unknown type, use mock client
		fmt.Println("Warning: Unknown registry type, using mock")
		return registry.NewMockClient()
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
