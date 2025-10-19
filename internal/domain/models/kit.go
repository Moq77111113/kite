package models

type Kit struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	Author      string     `json:"author"`
	Description string     `json:"description"`
	Files       []File     `json:"files"`
	Variables   []Variable `json:"variables"`
	Readme      string     `json:"readme"`
	Tags        []string   `json:"tags"`
}
