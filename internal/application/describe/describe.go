package describe

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/models"
	"github.com/moq77111113/kite/internal/domain/remote"
)

type Describe struct {
	repository    *remote.Repository
	installations *local.Tracker
}

func New(
	repository *remote.Repository,
	installations *local.Tracker,
) *Describe {
	return &Describe{
		repository:    repository,
		installations: installations,
	}
}

type Item struct {
	Name        string          `json:"name"`
	Version     string          `json:"version"`
	Description string          `json:"description"`
	Tags        []string        `json:"tags"`
	Installed   bool            `json:"installed"`
	Files       []models.File `json:"files"`
	Readme      string          `json:"readme"`
	Author      string          `json:"author"`
}

func (s *Describe) Execute(name string) (Item, error) {
	if name == "" {
		return Item{}, fmt.Errorf("name is required")
	}
	detail, err := s.repository.GetKit(name)
	if err != nil {
		if err := s.repository.Sync(); err != nil {
			return Item{}, err
		}
		detail, err = s.repository.GetKit(name)
	}
	if err != nil {
		return Item{}, err
	}

	installedMap := map[string]bool{}
	for _, kit := range s.installations.ListInstalled() {
		installedMap[kit.Name] = true
	}

	return Item{
		Name:        detail.Name,
		Version:     detail.Version,
		Description: detail.Description,
		Tags:        detail.Tags,
		Installed:   installedMap[detail.Name],
		Files:       detail.Files,
		Readme:      detail.Readme,
		Author:      detail.Author,
	}, nil
}
