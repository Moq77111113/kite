package parse

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/models"
	"gopkg.in/yaml.v3"
)

func ParseMetadata(yamlContent []byte) (*models.Metadata, error) {
	var metadata models.Metadata

	if err := yaml.Unmarshal(yamlContent, &metadata); err != nil {
		return nil, fmt.Errorf("invalid yaml: %w", err)
	}

	if err := ValidateMetadata(&metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

func ValidateMetadata(metadata *models.Metadata) error {
	if metadata.Name == "" {
		return fmt.Errorf("missing required field: name")
	}

	if metadata.Version == "" {
		return fmt.Errorf("missing required field: version")
	}

	return nil
}
