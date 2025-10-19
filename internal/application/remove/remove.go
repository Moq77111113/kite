package remove

import (
	"fmt"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/local"
)

type Remove struct {
	setupService *local.Manager
}

func New(setupService *local.Manager) *Remove {
	return &Remove{
		setupService: setupService,
	}
}

func (s *Remove) Execute(name, basePath string) error {
	kitPath := filepath.Join(basePath, name)

	if err := s.setupService.Uninstall(kitPath, name); err != nil {
		return fmt.Errorf("failed to uninstall kit: %w", err)
	}

	return nil
}
