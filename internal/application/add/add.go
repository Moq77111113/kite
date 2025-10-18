package add

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/install"
	"github.com/moq77111113/kite/internal/domain/repo"
)

type Add struct {
	setupService    *install.KitLifecycle
	conflictChecker *repo.ConflictChecker
	installations   *install.LocalKits
	repository      *repo.Repository
}

func New(
	setupService *install.KitLifecycle,
	conflictChecker *repo.ConflictChecker,
	installations *install.LocalKits,
	repository *repo.Repository,
) *Add {
	return &Add{
		setupService:    setupService,
		conflictChecker: conflictChecker,
		installations:   installations,
		repository:      repository,
	}
}

type Request struct {
	Name       string
	CustomPath string
	BasePath   string
}

type Result struct {
	Name          string
	Version       string
	InstalledPath string
	FilesCount    int
}

func (s *Add) Execute(req Request) (*Result, error) {
	destPath := s.setupService.CalculatePath(req.BasePath, req.Name, req.CustomPath)

	conflict, err := s.conflictChecker.Check(destPath)
	if err != nil {
		return nil, fmt.Errorf("failed to check for conflicts: %w", err)
	}

	if conflict.HasConflict {
		return nil, fmt.Errorf("conflict detected: %s at %s", conflict.Reason, conflict.ExistingPath)
	}

	kit, err := s.repository.GetKit(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch kit: %w", err)
	}

	if err := s.setupService.Install(kit, destPath); err != nil {
		return nil, fmt.Errorf("installation failed: %w", err)
	}

	return &Result{
		Name:          req.Name,
		Version:       kit.Version,
		InstalledPath: destPath,
		FilesCount:    len(kit.Files),
	}, nil
}
