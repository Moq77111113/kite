package template

import (
	"maps"

	"github.com/moq77111113/kite/internal/domain/models"
)

func buildLookup(values map[string]string, metadata []models.Variable) map[string]string {
	lookup := make(map[string]string)
	for _, v := range metadata {
		if v.Default != "" {
			lookup[v.Name] = v.Default
		}
	}
	maps.Copy(lookup, values)
	return lookup
}

func validateRequired(tokens []Token, lookup map[string]string, metadata []models.Variable) error {
	required := make(map[string]bool)
	for _, v := range metadata {
		if v.Required {
			required[v.Name] = true
		}
	}

	var missing []string
	for _, t := range tokens {
		if t.Type == TokenVariable && required[t.Value] {
			if _, ok := lookup[t.Value]; !ok {
				missing = append(missing, t.Value)
			}
		}
	}
	if len(missing) > 0 {
		return &MissingVariablesError{Variables: missing}
	}
	return nil
}
