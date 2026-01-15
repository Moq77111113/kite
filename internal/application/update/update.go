package update

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/local"
	"github.com/moq77111113/kite/internal/domain/remote"
)

type Update struct {
	repository      *remote.Repository
	tracker         *local.Tracker
	versionComp     *local.VersionComparator
	installer       *local.Installer
}

func New(
	repository *remote.Repository,
	tracker *local.Tracker,
	versionComp *local.VersionComparator,
	installer *local.Installer,
) *Update {
	return &Update{
		repository:      repository,
		tracker:         tracker,
		versionComp:     versionComp,
		installer:       installer,
	}
}

type Check struct {
	Name           string
	CurrentVersion string
	LatestVersion  string
}

func (s *Update) CheckAll() ([]Check, error) {
	var updates []Check
	installed := s.tracker.ListInstalled()

	for _, kit := range installed {
		latest, err := s.repository.GetKit(kit.ID)
		if err != nil {
			continue
		}

		hasUpdate, err := s.versionComp.IsUpdateAvailable(kit.Version, latest.Version)
		if err != nil || !hasUpdate {
			continue
		}

		updates = append(updates, Check{
			Name:           kit.ID,
			CurrentVersion: kit.Version,
			LatestVersion:  latest.Version,
		})
	}

	return updates, nil
}

func (s *Update) ApplyUpdate(name, basePath string) error {
	current, err := s.tracker.GetInstalled(name)
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

	destPath := s.installer.CalculatePath(basePath, name, "")
	if err := s.installer.Update(latest, destPath); err != nil {
		return fmt.Errorf("failed to update kit: %w", err)
	}

	return nil
}
