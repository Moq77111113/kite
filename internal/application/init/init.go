package init

import (
	"fmt"

	"github.com/moq77111113/kite/internal/infra/persistence/config"
)

type Request struct {
	Registry string
	Path     string
	Force    bool
}

type Result struct {
	Created  bool
	Registry string
	Path     string
}

type Init struct {
}

func New() *Init {
	return &Init{}
}

func (s *Init) Execute(req Request) (*Result, error) {
	existing, err := s.loadExisting(req.Force)
	if err != nil {
		return nil, err
	}

	registry := s.resolveRegistry(req.Registry, existing)
	if registry == "" {
		return nil, fmt.Errorf("registry is required")
	}

	path := s.resolvePath(req.Path, existing)

	created := existing == nil
	if err := s.saveConfig(existing, registry, path); err != nil {
		return nil, err
	}

	return &Result{
		Created:  created,
		Registry: registry,
		Path:     path,
	}, nil
}

func (s *Init) loadExisting(force bool) (*config.Config, error) {
	if !config.Exists("") {
		return nil, nil
	}

	cfg, err := config.Load("")
	if err != nil {
		return nil, fmt.Errorf("failed to load existing config: %w", err)
	}

	if !force {
		return nil, fmt.Errorf("config already exists (use --force to overwrite)")
	}

	return cfg, nil
}

func (s *Init) resolveRegistry(provided string, existing *config.Config) string {
	if provided != "" {
		return provided
	}
	if existing != nil {
		return existing.Registry
	}
	return ""
}

func (s *Init) resolvePath(provided string, existing *config.Config) string {
	if provided != "" {
		return provided
	}
	if existing != nil {
		return existing.Path
	}
	return config.DefaultPath
}

func (s *Init) saveConfig(existing *config.Config, registry, path string) error {
	if existing != nil {
		existing.Registry = registry
		existing.Path = path
		return config.Save(existing, "")
	}

	_, err := config.Init(registry, path, false, "")
	return err
}
