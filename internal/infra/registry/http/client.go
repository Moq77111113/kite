package http

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/moq77111113/kite/internal/domain/registry"
)

// HTTPClient is a registry API client using HTTP
type HTTPClient struct {
	baseURL string
	client  *resty.Client
}

// NewHTTPClient creates a new HTTP registry client
func NewHTTPClient(baseURL string) registry.Client {
	return &HTTPClient{
		baseURL: baseURL,
		client:  resty.New(),
	}
}

// ListTemplates fetches all available templates from the registry
func (c *HTTPClient) ListTemplates() ([]registry.TemplateSummary, error) {
	resp, err := c.client.R().
		SetResult(&registry.TemplateListResponse{}).
		Get(c.baseURL + "/templates")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch templates: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("registry returned status %d", resp.StatusCode())
	}

	result := resp.Result().(*registry.TemplateListResponse)
	return result.Templates, nil
}

// GetTemplate fetches a specific template by name
func (c *HTTPClient) GetTemplate(name string) (*registry.TemplateDetailResponse, error) {
	resp, err := c.client.R().
		SetResult(&registry.TemplateDetailResponse{}).
		Get(c.baseURL + "/templates/" + name)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch template %s: %w", name, err)
	}

	if resp.StatusCode() == 404 {
		return nil, fmt.Errorf("template %s not found", name)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("registry returned status %d", resp.StatusCode())
	}

	result := resp.Result().(*registry.TemplateDetailResponse)
	return result, nil
}
