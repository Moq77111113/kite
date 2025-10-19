package local

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Manager struct {
	installer     FsInstaller
	installations *Tracker
}

func NewManager(installer FsInstaller, installations *Tracker) *Manager {
	return &Manager{
		installer:     installer,
		installations: installations,
	}
}

func (s *Manager) Install(kit *models.Kit, destPath string) error {
	if kit == nil {
		return fmt.Errorf("kit cannot be nil")
	}

	if len(kit.Files) == 0 {
		return fmt.Errorf("kit %s has no files to install", kit.Name)
	}

	if destPath == "" {
		return fmt.Errorf("destination path cannot be empty")
	}

	if err := s.installer.Install(kit, destPath); err != nil {
		return fmt.Errorf("failed to install kit files: %w", err)
	}

	if err := s.installations.Record(kit.Name, kit.Version); err != nil {
		return fmt.Errorf("failed to record installation: %w", err)
	}

	return nil
}

func (s *Manager) Uninstall(kitPath, kitName string) error {
	if kitPath == "" {
		return fmt.Errorf("kit path cannot be empty")
	}

	absPath, err := filepath.Abs(kitPath)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("kit not found at %s", absPath)
	}

	if err := os.RemoveAll(absPath); err != nil {
		return fmt.Errorf("failed to remove kit directory: %w", err)
	}

	if err := s.installations.Unregister(kitName); err != nil {
		return fmt.Errorf("failed to unregister installation: %w", err)
	}

	return nil
}

func (s *Manager) Update(kit *models.Kit, destPath string) error {
	if err := s.installer.Install(kit, destPath); err != nil {
		return fmt.Errorf("failed to update kit files: %w", err)
	}

	if err := s.installations.UpdateVersion(kit.Name, kit.Version); err != nil {
		return fmt.Errorf("failed to update version: %w", err)
	}

	return nil
}

func (s *Manager) CalculatePath(basePath, kitName, customPath string) string {
	if customPath != "" {
		return customPath
	}
	return filepath.Join(basePath, kitName)
}
