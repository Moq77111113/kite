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

func (r *Tracker) Record(name, version string) error {
	if name == "" {
		return fmt.Errorf("kit name cannot be empty")
	}

	if version == "" {
		return fmt.Errorf("kit version cannot be empty")
	}

	return r.store.Add(name, version)
}

func (r *Tracker) Unregister(name string) error {
	if name == "" {
		return fmt.Errorf("kit name cannot be empty")
	}

	_, err := r.store.Get(name)
	if err != nil {
		return fmt.Errorf("kit %s is not installed", name)
	}

	return r.store.Remove(name)
}

func (r *Tracker) GetInstalled(name string) (*models.InstalledKit, error) {
	if name == "" {
		return nil, fmt.Errorf("kit name cannot be empty")
	}

	return r.store.Get(name)
}

func (r *Tracker) ListInstalled() []models.InstalledKit {
	return r.store.List()
}

func (r *Tracker) IsInstalled(name string) bool {
	_, err := r.store.Get(name)
	return err == nil
}

func (r *Tracker) UpdateVersion(name, version string) error {
	if !r.IsInstalled(name) {
		return fmt.Errorf("kit %s is not installed", name)
	}

	return r.Record(name, version)
}
