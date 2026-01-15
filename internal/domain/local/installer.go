package local

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Installer struct {
	writer  FileWriter
	tracker *Tracker
}

func NewInstaller(writer FileWriter, tracker *Tracker) *Installer {
	return &Installer{
		writer:  writer,
		tracker: tracker,
	}
}

func (i *Installer) Install(kit *models.Kit, destPath string) error {
	if kit == nil {
		return fmt.Errorf("kit cannot be nil")
	}

	if len(kit.Files) == 0 {
		return fmt.Errorf("kit %s has no files to install", kit.Name)
	}

	if destPath == "" {
		return fmt.Errorf("destination path cannot be empty")
	}

	if err := i.writer.Install(kit, destPath); err != nil {
		return fmt.Errorf("failed to install kit files: %w", err)
	}

	if err := i.tracker.Record(kit.ID, kit.Version); err != nil {
		return fmt.Errorf("failed to record installation: %w", err)
	}

	return nil
}

func (i *Installer) Uninstall(kitPath, kitID string) error {
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

	if err := i.tracker.Unregister(kitID); err != nil {
		return fmt.Errorf("failed to unregister installation: %w", err)
	}

	return nil
}

func (i *Installer) Update(kit *models.Kit, destPath string) error {
	if err := i.writer.Install(kit, destPath); err != nil {
		return fmt.Errorf("failed to update kit files: %w", err)
	}

	if err := i.tracker.UpdateVersion(kit.ID, kit.Version); err != nil {
		return fmt.Errorf("failed to update version: %w", err)
	}

	return nil
}

func (i *Installer) CalculatePath(basePath, kitName, customPath string) string {
	if customPath != "" {
		return customPath
	}
	return filepath.Join(basePath, kitName)
}
