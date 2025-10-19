package models

import "time"

type InstalledKit struct {
	Name      string
	Version   string
	Installed time.Time
}

type KitRegistry interface {
	Add(name, version string) error
	Remove(name string) error
	Get(name string) (*InstalledKit, error)
	List() []InstalledKit
}
