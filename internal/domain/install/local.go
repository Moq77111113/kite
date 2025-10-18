package install

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/port"
)

type LocalKits struct {
	store port.KitRegistry
}

func NewLocalKits(store port.KitRegistry) *LocalKits {
	return &LocalKits{
		store: store,
	}
}

func (r *LocalKits) Record(name, version string) error {
	if name == "" {
		return fmt.Errorf("kit name cannot be empty")
	}

	if version == "" {
		return fmt.Errorf("kit version cannot be empty")
	}

	return r.store.Add(name, version)
}

func (r *LocalKits) Unregister(name string) error {
	if name == "" {
		return fmt.Errorf("kit name cannot be empty")
	}

	_, err := r.store.Get(name)
	if err != nil {
		return fmt.Errorf("kit %s is not installed", name)
	}

	return r.store.Remove(name)
}

func (r *LocalKits) GetInstalled(name string) (*port.InstalledKit, error) {
	if name == "" {
		return nil, fmt.Errorf("kit name cannot be empty")
	}

	return r.store.Get(name)
}

func (r *LocalKits) ListInstalled() []port.InstalledKit {
	return r.store.List()
}

func (r *LocalKits) IsInstalled(name string) bool {
	_, err := r.store.Get(name)
	return err == nil
}

func (r *LocalKits) UpdateVersion(name, version string) error {
	if !r.IsInstalled(name) {
		return fmt.Errorf("kit %s is not installed", name)
	}

	return r.Record(name, version)
}
