package add

import (
	"fmt"

	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/pkg/console"
)

func addTemplate(svc *template.Service, name string) error {
	destPath := svc.GetDefaultPath(name)

	hasConflict, err := svc.CheckConflict(name)
	if err != nil {
		return fmt.Errorf("failed to check for conflicts: %w", err)
	}

	customPath := ""
	if hasConflict {
		action, err := promptConflictResolution(name, destPath)
		if err != nil {
			return err
		}

		switch action {
		case "cancel":
			return fmt.Errorf("installation cancelled")
		case "override":
		case "custom":
			customPath, err = promptCustomPath(svc.Config().Path, name)
			if err != nil {
				return err
			}
			destPath = customPath
		}
	}

	var tmpl *registry.TemplateDetailResponse
	err = console.Spinner(fmt.Sprintf("Fetching %s from registry", console.Cyan(name)), func() error {
		var fetchErr error
		tmpl, fetchErr = svc.FetchTemplate(name)
		return fetchErr
	})
	if err != nil {
		return err
	}

	err = console.Spinner(fmt.Sprintf("Installing %s", console.Cyan(name)), func() error {
		return svc.InstallTemplate(name, customPath, tmpl)
	})
	if err != nil {
		return err
	}

	console.Print("  %s %s → %s\n", console.Green("✓"), console.Bold(name), destPath)

	return nil
}
