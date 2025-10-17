package add

import (
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/moq77111113/kite/pkg/console"
)

// promptConflictResolution asks the user what to do when a directory already exists
func promptConflictResolution(name, existingPath string) (string, error) {
	console.Print("  %s Directory %s already exists\n", console.Yellow("⚠"), console.Cyan(existingPath))

	var action string
	prompt := &survey.Select{
		Message: "What would you like to do?",
		Options: []string{
			"Override (replace existing files)",
			"Use different path",
			"Cancel",
		},
	}

	if err := survey.AskOne(prompt, &action); err != nil {
		return "", err
	}

	switch action {
	case "Override (replace existing files)":
		return "override", nil
	case "Use different path":
		return "custom", nil
	default:
		return "cancel", nil
	}
}

// promptCustomPath asks the user for a custom installation path
func promptCustomPath(basePath, templateName string) (string, error) {
	var customName string
	prompt := &survey.Input{
		Message: "Enter custom path (relative to project):",
		Default: filepath.Join(basePath, templateName+"-2"),
		Help:    "The template will be installed at this location",
	}

	if err := survey.AskOne(prompt, &customName); err != nil {
		return "", err
	}

	return customName, nil
}
