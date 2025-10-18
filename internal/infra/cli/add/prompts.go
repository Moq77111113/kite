package add

import (
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/moq77111113/kite/pkg/console"
)

type ConflictAction string

const (
	ConflictOverride ConflictAction = "override"
	ConflictCustom   ConflictAction = "custom"
	ConflictCancel   ConflictAction = "cancel"
)

func promptConflictResolution(name, existingPath string) (ConflictAction, error) {
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
		return ConflictOverride, nil
	case "Use different path":
		return ConflictCustom, nil
	default:
		return ConflictCancel, nil
	}
}

func promptCustomPath(basePath, kitName string) (string, error) {
	var customName string
	prompt := &survey.Input{
		Message: "Enter custom path (relative to project):",
		Default: filepath.Join(basePath, kitName+"-2"),
		Help:    "The kit will be installed at this location",
	}

	if err := survey.AskOne(prompt, &customName); err != nil {
		return "", err
	}

	return customName, nil
}
