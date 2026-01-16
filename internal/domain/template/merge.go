package template

import "github.com/moq77111113/kite/internal/domain/models"

func MergeWithMetadata(detected []string, metadata []models.Variable) []models.Variable {
	metadataMap := make(map[string]models.Variable)
	for _, v := range metadata {
		metadataMap[v.Name] = v
	}

	var result []models.Variable
	for _, name := range detected {
		if meta, ok := metadataMap[name]; ok {
			result = append(result, meta)
		} else {
			result = append(result, models.Variable{Name: name, Type: "string"})
		}
	}
	return result
}
