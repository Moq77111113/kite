package add

import (
	"fmt"
	"os"

	"github.com/moq77111113/kite/internal/application/add"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/pkg/console"
)

func addKit(addSvc *add.Add, cfg *config.Config, name string) error {

	destPath := cfg.Path + "/" + name

	hasConflict, err := checkConflict(destPath)
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
		case ConflictCancel:
			return fmt.Errorf("installation cancelled")
		case ConflictOverride:
		case ConflictCustom:
			customPath, err = promptCustomPath(cfg.Path, name)
			if err != nil {
				return err
			}
			destPath = customPath
		}
	}

	var result *add.Result
	err = console.Spinner(fmt.Sprintf("Installing %s", console.Cyan(name)), func() error {
		var execErr error
		result, execErr = addSvc.Execute(add.Request{
			Name:       name,
			CustomPath: customPath,
			BasePath:   cfg.Path,
		})
		return execErr
	})
	if err != nil {
		return err
	}

	console.Print("  %s %s %s → %s\n",
		console.Green("✓"),
		console.Bold(result.Name),
		console.Dim(result.Version),
		result.InstalledPath,
	)

	return nil
}

func checkConflict(destPath string) (bool, error) {
	return pathExists(destPath), nil
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
