package template

import (
	"fmt"
	"time"

	"github.com/moq77111113/kite/internal/application/registry"
	reg "github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/domain/template"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

// Service handles template operations orchestration
type Service struct {
	config  *config.Config
	client  reg.Client
	manager *template.Manager
}

// NewService creates a new template service
func NewService(cfg *config.Config) *Service {
	client := registry.NewClient(cfg.Registry)
	manager := template.NewManager(cfg, client)

	return &Service{
		config:  cfg,
		client:  client,
		manager: manager,
	}
}

// Add adds a template to the project
func (s *Service) Add(name string) error {
	return s.manager.Add(name)
}

// Remove removes a template from the project
func (s *Service) Remove(name string) error {
	return s.manager.Remove(name)
}

// CheckUpdate checks if an update is available for a template
func (s *Service) CheckUpdate(name string) (*template.UpdateInfo, error) {
	return s.manager.CheckUpdate(name)
}

// GetInstalled gets an installed template's info
func (s *Service) GetInstalled(name string) (*reg.InstalledTemplate, error) {
	tmpl, exists := s.config.GetTemplate(name)
	if !exists {
		return nil, fmt.Errorf("template %s is not installed", name)
	}
	return &reg.InstalledTemplate{
		Name:      name,
		Version:   tmpl.Version,
		Installed: convertTimestamp(tmpl.Installed),
	}, nil
}

// ListInstalled lists all installed templates
func (s *Service) ListInstalled() []reg.InstalledTemplate {
	var installed []reg.InstalledTemplate
	for name, tmpl := range s.config.Templates {
		installed = append(installed, reg.InstalledTemplate{
			Name:      name,
			Version:   tmpl.Version,
			Installed: convertTimestamp(tmpl.Installed),
		})
	}
	return installed
}

// ListAvailable lists all templates from the registry
func (s *Service) ListAvailable() ([]reg.TemplateSummary, error) {
	return s.client.ListTemplates()
}

// GetDetails gets template details from registry
func (s *Service) GetDetails(name string) (*reg.TemplateDetailResponse, error) {
	return s.client.GetTemplate(name)
}

// Config returns the current config
func (s *Service) Config() *config.Config {
	return s.config
}

// convertTimestamp converts Unix timestamp to time.Time
func convertTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}
