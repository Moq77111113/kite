package list

import (
	"fmt"
	"time"

	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/models"
	"github.com/moq77111113/kite/internal/domain/remote"
)

type List struct {
	repository *remote.Repository
	tracker    *local.Tracker
}

func New(
	repository *remote.Repository,
	tracker *local.Tracker,
) *List {
	return &List{
		repository: repository,
		tracker:    tracker,
	}
}

type Item struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Author      string     `json:"author"`
	Installed   bool       `json:"installed"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
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

func (s *List) enrichWithInstallationStatus(available []models.KitSummary) []Item {
	installed := s.tracker.ListInstalled()

	installedMap := make(map[string]bool)
	for _, kit := range installed {
		installedMap[kit.Name] = true
	}

	var items []Item
	for _, kit := range available {
		item := Item{
			Name:        kit.Name,
			Version:     kit.Version,
			Description: kit.Description,
			Tags:        kit.Tags,
			Installed:   installedMap[kit.Name],
			Author:      kit.Author,
			LastUpdated: kit.LastUpdated,
		}
		if item.Tags == nil {
			item.Tags = []string{}
		}
		items = append(items, item)
	}

	return items
}
