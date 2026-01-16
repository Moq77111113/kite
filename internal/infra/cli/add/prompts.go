package add

import (
	"fmt"
	"maps"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/moq77111113/kite/internal/domain/models"
	"github.com/moq77111113/kite/pkg/console"
)

type ConflictAction string

const (
	ConflictOverride ConflictAction = "override"
	ConflictCustom   ConflictAction = "custom"
	ConflictCancel   ConflictAction = "cancel"
)

func promptConflictResolution(_, existingPath string) (ConflictAction, error) {
	console.Print("  %s Directory %s already exists\n", console.Yellow("⚠"), console.Cyan(existingPath))
	var action string
	prompt := &survey.Select{
		Message: "What would you like to do?",
		Options: []string{"Override (replace existing files)", "Use different path", "Cancel"},
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
	}
	if err := survey.AskOne(prompt, &customName); err != nil {
		return "", err
	}
	return customName, nil
}

func collectMissingVariables(defs []models.Variable, provided map[string]string) (map[string]string, error) {
	result := make(map[string]string)
	maps.Copy(result, provided)

	var questions []*survey.Question
	for _, v := range defs {
		if _, ok := result[v.Name]; ok {
			continue
		}
		msg := v.Name
		if v.Description != "" {
			msg = fmt.Sprintf("%s (%s)", v.Name, v.Description)
		}
		q := &survey.Question{
			Name:   v.Name,
			Prompt: &survey.Input{Message: msg, Default: v.Default},
		}
		questions = append(questions, q)
	}

	if len(questions) == 0 {
		return result, nil
	}

	console.Print("  %s Configure variables:\n", console.Cyan("?"))
	answers := make(map[string]any)
	if err := survey.Ask(questions, &answers); err != nil {
		return nil, err
	}
	for k, v := range answers {
		result[k] = fmt.Sprintf("%v", v)
	}
	return result, nil
}
