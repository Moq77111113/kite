package repo

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/port"
	"github.com/moq77111113/kite/internal/domain/scan"
	registry "github.com/moq77111113/kite/internal/domain/types"
)

type Repository struct {
	storage port.Storage
}

func NewRepository(store port.Storage) *Repository {
	return &Repository{
		storage: store,
	}
}

func (r *Repository) GetKit(name string) (*registry.KitDetailResponse, error) {
	if name == "" {
		return nil, fmt.Errorf("kit name cannot be empty")
	}

	kit, err := scan.FindKit(r.storage, name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch kit %s: %w", name, err)
	}

	return kit, nil
}

func (r *Repository) ListAvailable() ([]registry.KitSummary, error) {
	kits, err := scan.ScanForKits(r.storage)
	if err != nil {
		return nil, fmt.Errorf("failed to list kits: %w", err)
	}

	return kits, nil
}

func (r *Repository) Sync() error {
	if err := r.storage.Sync(); err != nil {
		return fmt.Errorf("failed to sync: %w", err)
	}
	return nil
}
