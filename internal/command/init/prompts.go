package initcmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// RegistryType represents a type of registry
type RegistryType string

const (
	RegistryTypeGit   RegistryType = "git"
	RegistryTypeLocal RegistryType = "local"
	RegistryTypeHTTP  RegistryType = "http"
)

// Registry selection options
var registryOptions = map[RegistryType]string{
	RegistryTypeGit:   "Git repository (GitHub, GitLab, etc.)",
	RegistryTypeLocal: "Local directory (file system)",
	RegistryTypeHTTP:  "HTTP API (custom server)",
}

// AskRegistryType prompts user to select a registry type
func AskRegistryType() (RegistryType, error) {
	options := []string{
		registryOptions[RegistryTypeGit],
		registryOptions[RegistryTypeLocal],
		registryOptions[RegistryTypeHTTP],
	}

	var selected string
	prompt := &survey.Select{
		Message: "What type of registry do you want to use?",
		Options: options,
		Default: registryOptions[RegistryTypeGit],
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return "", err
	}

	return getRegistryType(selected), nil
}

// AskRegistryURL prompts user for registry URL based on type
func AskRegistryURL(registryType RegistryType) (string, error) {
	prompts := map[RegistryType]*survey.Input{
		RegistryTypeGit: {
			Message: "Registry Git URL:",
			Help:    "Examples: https://github.com/your-org/kite-registry, git@github.com:your-org/kite-registry.git",
		},
		RegistryTypeLocal: {
			Message: "Registry local path:",
			Help:    "Examples: ./my-templates, /path/to/templates",
		},
		RegistryTypeHTTP: {
			Message: "Registry HTTP URL:",
			Help:    "Example: https://api.kite.sh",
		},
	}

	prompt, ok := prompts[registryType]
	if !ok {
		prompt = &survey.Input{Message: "Registry URL:"}
	}

	var url string
	if err := survey.AskOne(prompt, &url); err != nil {
		return "", err
	}

	url = strings.TrimSpace(url)
	if url == "" {
		return "", fmt.Errorf("registry URL cannot be empty")
	}

	return url, nil
}

// AskPath prompts user for installation path
func AskPath(defaultPath string) (string, error) {
	prompt := &survey.Input{
		Message: "Where should templates be installed?",
		Default: defaultPath,
		Help:    "Relative path from project root",
	}

	var path string
	if err := survey.AskOne(prompt, &path); err != nil {
		return "", err
	}

	return path, nil
}

// AskConfirm prompts user for yes/no confirmation
func AskConfirm(message string, defaultValue bool) (bool, error) {
	var result bool
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultValue,
	}

	if err := survey.AskOne(prompt, &result); err != nil {
		return false, err
	}

	return result, nil
}

func getRegistryType(display string) RegistryType {
	for typ, disp := range registryOptions {
		if disp == display {
			return typ
		}
	}
	return RegistryTypeGit
}
