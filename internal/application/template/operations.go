package template

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/moq77111113/kite/internal/application/registry"
	reg "github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/domain/template"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

// Service handles template operations orchestration
type Service struct {
	config     *config.Config
	client     reg.Client
	repository *template.Repository
}

// NewService creates a new template service
func NewService(cfg *config.Config) *Service {
	client := registry.NewClient(cfg.Registry)
	repository := template.NewRepository(cfg, client)

	return &Service{
		config:     cfg,
		client:     client,
		repository: repository,
	}
}

// Add adds a template to the project
func (s *Service) Add(name, customPath string) error {
	return s.repository.Add(name, customPath)
}

// CheckConflict checks if a template directory already exists
func (s *Service) CheckConflict(name string) (bool, error) {
	destPath := filepath.Join(s.config.Path, name)
	return s.repository.CheckConflict(destPath)
}

// GetDefaultPath returns the default installation path for a template
func (s *Service) GetDefaultPath(name string) string {
	return filepath.Join(s.config.Path, name)
}

// Remove removes a template from the project
func (s *Service) Remove(name string) error {
	return s.repository.Remove(name)
}

// CheckUpdate checks if an update is available for a template
func (s *Service) CheckUpdate(name string) (*template.UpdateInfo, error) {
	return s.repository.CheckUpdate(name)
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
