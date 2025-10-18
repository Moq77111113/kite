package models

import registry "github.com/moq77111113/kite/internal/domain/types"

type Metadata struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Version     string              `yaml:"version"`
	Author      string              `yaml:"author"`
	Tags        []string            `yaml:"tags"`
	Variables   []registry.Variable `yaml:"variables,omitempty"`
}

func (m *Metadata) ToKitSummary() registry.KitSummary {
	return registry.KitSummary{
		Name:        m.Name,
		Description: m.Description,
		Version:     m.Version,
		Tags:        m.Tags,
		Author:      m.Author,
	}
}

func (m *Metadata) ToKitDetail(files []registry.KitFile, readme string) *registry.KitDetailResponse {
	return &registry.KitDetailResponse{
		Name:        m.Name,
		Version:     m.Version,
		Author:      m.Author,
		Description: m.Description,
		Files:       files,
		Variables:   m.Variables,
		Readme:      readme,
		Tags:        m.Tags,
	}
}
