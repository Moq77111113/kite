package models

import "time"

type File struct {
	Path        string     `json:"path"`
	Content     string     `json:"content"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}
