package local

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/models"
)

type Tracker struct {
	store models.KitRegistry
}

func NewTracker(store models.KitRegistry) *Tracker {
	return &Tracker{
		store: store,
	}
}

func (r *Tracker) Record(id, version string) error {
	if id == "" {
		return fmt.Errorf("kit id cannot be empty")
	}

	if version == "" {
		return fmt.Errorf("kit version cannot be empty")
	}

	return r.store.Add(id, version)
}

func (r *Tracker) Unregister(id string) error {
	if id == "" {
		return fmt.Errorf("kit id cannot be empty")
	}

	_, err := r.store.Get(id)
	if err != nil {
		return fmt.Errorf("kit %s is not installed", id)
	}

	return r.store.Remove(id)
}

func (r *Tracker) GetInstalled(id string) (*models.InstalledKit, error) {
	if id == "" {
		return nil, fmt.Errorf("kit id cannot be empty")
	}

	return r.store.Get(id)
}

func (r *Tracker) ListInstalled() []models.InstalledKit {
	return r.store.List()
}

func (r *Tracker) IsInstalled(id string) bool {
	_, err := r.store.Get(id)
	return err == nil
}

func (r *Tracker) UpdateVersion(id, version string) error {
	if !r.IsInstalled(id) {
		return fmt.Errorf("kit %s is not installed", id)
	}

	return r.Record(id, version)
}
