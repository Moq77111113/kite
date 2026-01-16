package add

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func parseVariables(flags []string, filePath string) (map[string]string, error) {
	vars := make(map[string]string)

	if filePath != "" {
		fileVars, err := loadVariableFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("loading var-file: %w", err)
		}
		for k, v := range fileVars {
			vars[k] = v
		}
	}

	for _, flag := range flags {
		key, value, err := parseVarFlag(flag)
		if err != nil {
			return nil, err
		}
		vars[key] = value
	}
	return vars, nil
}

func parseVarFlag(flag string) (string, string, error) {
	idx := strings.Index(flag, "=")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid variable format '%s': expected key=value", flag)
	}
	key := strings.TrimSpace(flag[:idx])
	if key == "" {
		return "", "", fmt.Errorf("variable name cannot be empty in '%s'", flag)
	}
	return key, flag[idx+1:], nil
}

func loadVariableFile(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nested struct {
		Variables map[string]any `yaml:"variables"`
	}
	if err := yaml.Unmarshal(data, &nested); err == nil && nested.Variables != nil {
		return convertToStringMap(nested.Variables), nil
	}

	var flat map[string]any
	if err := yaml.Unmarshal(data, &flat); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	return convertToStringMap(flat), nil
}

func convertToStringMap(m map[string]any) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = fmt.Sprintf("%v", v)
	}
	return result
}
