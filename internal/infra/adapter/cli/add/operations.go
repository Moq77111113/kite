package add

import (
	"fmt"

	"github.com/moq77111113/kite/internal/application/template"
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

	var installErr error
	err = console.Spinner(fmt.Sprintf("Fetching %s", console.Cyan(name)), func() error {
		installErr = svc.Add(name, customPath)
		return installErr
	})

	if err != nil || installErr != nil {
		if installErr != nil {
			return installErr
		}
		return err
	}

	console.Print("  %s %s → %s\n", console.Green("✓"), console.Bold(name), destPath)

	return nil
}
