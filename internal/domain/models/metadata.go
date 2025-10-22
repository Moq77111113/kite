package models

import "time"

type Metadata struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Version     string     `yaml:"version"`
	Tags        []string   `yaml:"tags"`
	Author      string     `yaml:"author"`
	Variables   []Variable `yaml:"variables,omitempty"`
}

func (m *Metadata) ToKitSummary(lastUpdated *time.Time) KitSummary {
	return KitSummary{
		Name:        m.Name,
		Description: m.Description,
		Version:     m.Version,
		Tags:        m.Tags,
		Author:      m.Author,
		LastUpdated: lastUpdated,
	}
}

func (m *Metadata) ToKitDetail(files []File, readme string) *Kit {
	return &Kit{
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
