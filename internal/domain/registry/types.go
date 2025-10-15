package registry

import "time"

// TemplateListResponse is the response from GET /templates
type TemplateListResponse struct {
	Templates []TemplateSummary `json:"templates"`
}

// TemplateSummary represents a template in the list view
type TemplateSummary struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`

	Tags   []string `json:"tags"`
	Author string   `json:"author"`
}

// TemplateDetailResponse is the response from GET /templates/{name}
type TemplateDetailResponse struct {
	Name        string         `json:"name"`
	Version     string         `json:"version"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	Files       []TemplateFile `json:"files"`
	Variables   []Variable     `json:"variables"`
	Readme      string         `json:"readme"`
}

// TemplateFile represents a file in a template
type TemplateFile struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

// Variable represents a template variable
type Variable struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// InstalledTemplate represents a template installed locally
type InstalledTemplate struct {
	Name      string
	Version   string
	Installed time.Time
}
