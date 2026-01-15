package models

import "time"

type InstalledKit struct {
	ID        string
	Version   string
	Installed time.Time
}

type KitRegistry interface {
	Add(id, version string) error
	Remove(id string) error
	Get(id string) (*InstalledKit, error)
	List() []InstalledKit
}
