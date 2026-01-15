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

func (m *Metadata) ToKitSummary(dirName string, lastUpdated *time.Time) KitSummary {
	name := m.Name
	if name == "" {
		name = dirName
	}
	return KitSummary{
		ID:          dirName,
		Name:        name,
		Description: m.Description,
		Version:     m.Version,
		Tags:        m.Tags,
		Author:      m.Author,
		LastUpdated: lastUpdated,
	}
}

func (m *Metadata) ToKitDetail(dirName string, files []File, readme string) *Kit {
	name := m.Name
	if name == "" {
		name = dirName
	}
	return &Kit{
		ID:          dirName,
		Name:        name,
		Version:     m.Version,
		Author:      m.Author,
		Description: m.Description,
		Files:       files,
		Variables:   m.Variables,
		Readme:      readme,
		Tags:        m.Tags,
	}
}
