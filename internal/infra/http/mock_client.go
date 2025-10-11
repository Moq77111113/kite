package http

import (
	"fmt"

	"github.com/moq77111113/kite/api/registry/v1"
)

// MockClient is a mock registry client for testing
type MockClient struct {
	templates []registry.TemplateSummary
	details   map[string]*registry.TemplateDetailResponse
}

// NewMockClient creates a mock client with example templates
func NewMockClient() *MockClient {
	return &MockClient{
		templates: mockTemplates(),
		details:   mockDetails(),
	}
}

func (m *MockClient) ListTemplates() ([]registry.TemplateSummary, error) {
	return m.templates, nil
}

func (m *MockClient) GetTemplate(name string) (*registry.TemplateDetailResponse, error) {
	detail, ok := m.details[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}
	return detail, nil
}

func mockTemplates() []registry.TemplateSummary {
	return []registry.TemplateSummary{
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
	}
}

func mockDetails() map[string]*registry.TemplateDetailResponse {
	return map[string]*registry.TemplateDetailResponse{
		"aws-vpc": {
			Name:        "aws-vpc",
			Version:     "1.0.0",
			Description: "Production-ready AWS VPC",
			Readme:      "# AWS VPC Template\n\nA production-ready VPC setup for AWS.",
			Files: []registry.TemplateFile{
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
			Files: []registry.TemplateFile{
				{Path: "prometheus.yaml", Content: "# Prometheus config\n"},
				{Path: "grafana.yaml", Content: "# Grafana config\n"},
			},
		},
		"docker-redis": {
			Name:        "docker-redis",
			Version:     "1.2.0",
			Description: "Redis with persistence",
			Readme:      "# Docker Redis\n\nRedis with volume persistence.",
			Files: []registry.TemplateFile{
				{Path: "docker-compose.yml", Content: "# Redis compose\n"},
				{Path: "redis.conf", Content: "# Redis config\n"},
			},
		},
	}
}
