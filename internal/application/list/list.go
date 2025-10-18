package list

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/install"
	"github.com/moq77111113/kite/internal/domain/repo"
	registry "github.com/moq77111113/kite/internal/domain/types"
)

type List struct {
	repository    *repo.Repository
	installations *install.LocalKits
}

func New(
	repository *repo.Repository,
	installations *install.LocalKits,
) *List {
	return &List{
		repository:    repository,
		installations: installations,
	}
}

type Item struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Installed   bool     `json:"installed"`
}

func (s *List) Execute() ([]Item, error) {
	available, err := s.repository.ListAvailable()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch available kits: %w", err)
	}

	if len(available) == 0 {
		if err := s.repository.Sync(); err != nil {
			return s.enrichWithInstallationStatus(available), nil
		}

		available, err = s.repository.ListAvailable()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch kits after sync: %w", err)
		}
	}

	return s.enrichWithInstallationStatus(available), nil
}

func (s *List) enrichWithInstallationStatus(available []registry.KitSummary) []Item {
	installed := s.installations.ListInstalled()

	installedMap := make(map[string]bool)
	for _, kit := range installed {
		installedMap[kit.Name] = true
	}

	var items []Item
	for _, kit := range available {
		items = append(items, Item{
			Name:        kit.Name,
			Version:     kit.Version,
			Description: kit.Description,
			Tags:        kit.Tags,
			Installed:   installedMap[kit.Name],
		})
	}

	return items
}
