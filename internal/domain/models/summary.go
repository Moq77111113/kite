package models

import "time"

type KitListResponse struct {
	Kits []KitSummary `json:"kits"`
}

type KitSummary struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Version     string     `json:"version"`
	Tags        []string   `json:"tags"`
	Author      string     `json:"author"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}
