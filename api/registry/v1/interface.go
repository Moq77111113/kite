package registry

// Client defines the interface for registry operations
type Client interface {
	ListTemplates() ([]TemplateSummary, error)
	GetTemplate(name string) (*TemplateDetailResponse, error)
}
