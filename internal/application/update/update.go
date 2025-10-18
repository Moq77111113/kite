package update

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/install"
	"github.com/moq77111113/kite/internal/domain/repo"
)

type Update struct {
	repository    *repo.Repository
	installations *install.LocalKits
	versionComp   *repo.VersionComparator
	setupService  *install.KitLifecycle
}

func New(
	repository *repo.Repository,
	installations *install.LocalKits,
	versionComp *repo.VersionComparator,
	setupService *install.KitLifecycle,
) *Update {
	return &Update{
		repository:    repository,
		installations: installations,
		versionComp:   versionComp,
		setupService:  setupService,
	}
}

type Check struct {
	Name           string
	CurrentVersion string
	LatestVersion  string
}

func (s *Update) CheckAll() ([]Check, error) {
	var updates []Check
	installed := s.installations.ListInstalled()

	for _, kit := range installed {
		latest, err := s.repository.GetKit(kit.Name)
		if err != nil {
			continue
		}

		hasUpdate, err := s.versionComp.IsUpdateAvailable(kit.Version, latest.Version)
		if err != nil || !hasUpdate {
			continue
		}

		updates = append(updates, Check{
			Name:           kit.Name,
			CurrentVersion: kit.Version,
			LatestVersion:  latest.Version,
		})
	}

	return updates, nil
}

func (s *Update) ApplyUpdate(name, basePath string) error {
	current, err := s.installations.GetInstalled(name)
	if err != nil {
		return fmt.Errorf("kit not installed: %w", err)
	}

	latest, err := s.repository.GetKit(name)
	if err != nil {
		return fmt.Errorf("failed to fetch latest version: %w", err)
	}

	hasUpdate, err := s.versionComp.IsUpdateAvailable(current.Version, latest.Version)
	if err != nil {
		return fmt.Errorf("version comparison failed: %w", err)
	}

	if !hasUpdate {
		return fmt.Errorf("no update available (current: %s, latest: %s)", current.Version, latest.Version)
	}

	destPath := s.setupService.CalculatePath(basePath, name, "")
	if err := s.setupService.Update(latest, destPath); err != nil {
		return fmt.Errorf("failed to update kit: %w", err)
	}

	return nil
}
