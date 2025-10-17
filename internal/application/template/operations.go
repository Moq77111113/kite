package template

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/moq77111113/kite/internal/application/registry"
	reg "github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/domain/template"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/internal/infra/registry/git"
)

type Service struct {
	config     *config.Config
	client     reg.Client
	repository *template.Repository
}

// NewService creates a new template service with optional sync callback
func NewService(cfg *config.Config, syncCallback git.SyncCallback) *Service {
	client := registry.NewClient(cfg.Registry, syncCallback)
	repository := template.NewRepository(cfg, client)

	return &Service{
		config:     cfg,
		client:     client,
		repository: repository,
	}
}

func (s *Service) Add(name, customPath string) error {
	return s.repository.Add(name, customPath)
}

func (s *Service) FetchTemplate(name string) (*reg.TemplateDetailResponse, error) {
	return s.repository.FetchTemplate(name)
}

func (s *Service) InstallTemplate(name, customPath string, template *reg.TemplateDetailResponse) error {
	return s.repository.InstallTemplate(name, customPath, template)
}

func (s *Service) CheckConflict(name string) (bool, error) {
	destPath := filepath.Join(s.config.Path, name)
	return s.repository.CheckConflict(destPath)
}

func (s *Service) GetDefaultPath(name string) string {
	return filepath.Join(s.config.Path, name)
}

func (s *Service) Remove(name string) error {
	return s.repository.Remove(name)
}

func (s *Service) CheckUpdate(name string) (*template.UpdateInfo, error) {
	return s.repository.CheckUpdate(name)
}

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

func (s *Service) ListAvailable() ([]reg.TemplateSummary, error) {
	return s.client.ListTemplates()
}

func (s *Service) GetDetails(name string) (*reg.TemplateDetailResponse, error) {
	return s.client.GetTemplate(name)
}

func (s *Service) Config() *config.Config {
	return s.config
}

func convertTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}
