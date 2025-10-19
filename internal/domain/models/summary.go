package models

type KitListResponse struct {
	Kits []KitSummary `json:"kits"`
}

type KitSummary struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author"`
}
