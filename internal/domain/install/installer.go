package install

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/types"
)

type FsInstaller interface {
	Install(kit *types.KitDetailResponse, destPath string) error
}

type installer struct{}

func NewFsInstaller() FsInstaller {
	return &installer{}
}

func (i *installer) Install(kit *types.KitDetailResponse, destPath string) error {
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for _, file := range kit.Files {
		if err := i.writeFile(destPath, file); err != nil {
			return err
		}
	}

	return nil
}

func (i *installer) writeFile(basePath string, file types.KitFile) error {
	filePath := filepath.Join(basePath, file.Path)
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	if err := os.WriteFile(filePath, []byte(file.Content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}
