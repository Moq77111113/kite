package registry

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// HTTPClient is a registry API client using HTTP
type HTTPClient struct {
	baseURL string
	client  *resty.Client
}

// NewHTTPClient creates a new HTTP registry client
func NewHTTPClient(baseURL string) Client {
	return &HTTPClient{
		baseURL: baseURL,
		client:  resty.New(),
	}
}

// ListTemplates fetches all available templates from the registry
func (c *HTTPClient) ListTemplates() ([]TemplateSummary, error) {
	resp, err := c.client.R().
		SetResult(&TemplateListResponse{}).
		Get(c.baseURL + "/templates")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch templates: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("registry returned status %d", resp.StatusCode())
	}

	result := resp.Result().(*TemplateListResponse)
	return result.Templates, nil
}

// GetTemplate fetches a specific template by name
func (c *HTTPClient) GetTemplate(name string) (*TemplateDetailResponse, error) {
	resp, err := c.client.R().
		SetResult(&TemplateDetailResponse{}).
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

	result := resp.Result().(*TemplateDetailResponse)
	return result, nil
}

// MockClient is a mock registry client for testing
type MockClient struct {
	templates []TemplateSummary
	details   map[string]*TemplateDetailResponse
}

// NewMockClient creates a mock client with example templates
func NewMockClient() *MockClient {
	return &MockClient{
		templates: []TemplateSummary{
			{
				Name:        "aws-vpc",
				Description: "Production-ready AWS VPC",
				Version:     "1.0.0",
				Tags:        []string{"aws", "networking"},
				Author:      "Kite Team",
			},
			{
				Name:        "k8s-monitoring",
				Description: "Prometheus + Grafana stack",
				Version:     "2.1.0",
				Tags:        []string{"kubernetes", "monitoring"},
				Author:      "Kite Team",
			},
			{
				Name:        "docker-redis",
				Description: "Redis with persistence",
				Version:     "1.2.0",
				Tags:        []string{"docker", "database"},
				Author:      "Kite Team",
			},
		},
		details: map[string]*TemplateDetailResponse{
			"aws-vpc": {
				Name:        "aws-vpc",
				Version:     "1.0.0",
				Description: "Production-ready AWS VPC",
				Readme:      "# AWS VPC Template\n\nA production-ready VPC setup for AWS.",
				Files: []TemplateFile{
					{Path: "main.tf", Content: "# VPC Configuration\n"},
					{Path: "variables.tf", Content: "# Variables\n"},
					{Path: "outputs.tf", Content: "# Outputs\n"},
				},
			},
			"k8s-monitoring": {
				Name:        "k8s-monitoring",
				Version:     "2.1.0",
				Description: "Prometheus + Grafana stack",
				Readme:      "# Kubernetes Monitoring\n\nComplete monitoring stack.",
				Files: []TemplateFile{
					{Path: "prometheus.yaml", Content: "# Prometheus config\n"},
					{Path: "grafana.yaml", Content: "# Grafana config\n"},
				},
			},
			"docker-redis": {
				Name:        "docker-redis",
				Version:     "1.2.0",
				Description: "Redis with persistence",
				Readme:      "# Docker Redis\n\nRedis with volume persistence.",
				Files: []TemplateFile{
					{Path: "docker-compose.yml", Content: "# Redis compose\n"},
					{Path: "redis.conf", Content: "# Redis config\n"},
				},
			},
		},
	}
}

func (m *MockClient) ListTemplates() ([]TemplateSummary, error) {
	return m.templates, nil
}

func (m *MockClient) GetTemplate(name string) (*TemplateDetailResponse, error) {
	detail, ok := m.details[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}
	return detail, nil
}

// LoadFromFile loads templates from a local JSON file (for file-based registry)
func LoadFromFile(path string) ([]TemplateSummary, error) {
	// Read the file
	data, err := readFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read registry file: %w", err)
	}

	// Parse as template list
	var listResp TemplateListResponse
	if err := json.Unmarshal(data, &listResp); err != nil {
		return nil, fmt.Errorf("failed to parse registry file: %w", err)
	}

	return listResp.Templates, nil
}

// Helper to read file (can be mocked in tests)
var readFile = func(path string) ([]byte, error) {
	return nil, fmt.Errorf("not implemented - use os.ReadFile")
}
