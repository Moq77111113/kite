package remote

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Repository struct {
	storage Storage
}

func NewRepository(store Storage) *Repository {
	return &Repository{
		storage: store,
	}
}

func (r *Repository) GetKit(name string) (*models.Kit, error) {
	if name == "" {
		return nil, fmt.Errorf("kit name cannot be empty")
	}

	kit, err := FindKit(r.storage, name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch kit %s: %w", name, err)
	}

	return kit, nil
}

func (r *Repository) ListAvailable() ([]models.KitSummary, error) {
	kits, err := ScanForKits(r.storage)
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
