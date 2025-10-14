package add

import (
	"fmt"
	"path/filepath"

	"github.com/moq77111113/kite/internal/infra/adapter/cli/cmdutil"
	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add <template-name> [template-name...]",
		Short: "Download a template from registry and add to project",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runAdd,
	}
}

func runAdd(cmd *cobra.Command, args []string) error {
	cfg, err := cmdutil.LoadConfig(cmd)
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	s := template.NewService(cfg)

	console.EmptyLine()
	installed := 0
	for i, name := range args {
		if i > 0 {
			console.EmptyLine()
		}
		if err := addTemplate(s, name); err != nil {
			if err.Error() == fmt.Sprintf("template %s is already installed", name) {
				console.Print("  %s %s\n", console.Yellow("⚠"), console.Yellow("Already installed"))
			} else {
				console.Print("  %s %s\n", console.Red("✗"), console.Red("Failed: %v", err))
			}
			continue
		}
		installed++
	}

	if installed > 0 {
		console.EmptyLine()
		console.Success(fmt.Sprintf("Successfully installed %d template(s)", installed))
	}

	return nil
}

func addTemplate(svc *template.Service, name string) error {
	var installErr error

	err := console.Spinner(fmt.Sprintf("Fetching %s", console.Cyan(name)), func() error {
		installErr = svc.Add(name)
		return installErr
	})

	if err != nil || installErr != nil {
		if installErr != nil {
			return installErr
		}
		return err
	}

	cfg := svc.Config()
	destPath := filepath.Join(cfg.Path, name)
	console.Print("  %s %s → %s\n", console.Green("✓"), console.Bold(name), destPath)

	return nil
}
