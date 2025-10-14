package update

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	registryv1 "github.com/moq77111113/kite/api/registry/v1"
	"github.com/moq77111113/kite/internal/command/cmdutil"
	"github.com/moq77111113/kite/internal/domain/config"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update installed templates to latest versions",
		RunE:  runUpdate,
	}
}

func runUpdate(cmd *cobra.Command, args []string) error {
	cfg, err := cmdutil.LoadConfig(cmd)
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	if len(cfg.Templates) == 0 {
		console.EmptyLine()
		console.Println(console.Dim("No templates installed"))
		return nil
	}

	client := registry.NewClient(cfg.Registry)
	updates, err := checkForUpdates(cfg, client)
	if err != nil {
		return err
	}

	if len(updates) == 0 {
		console.EmptyLine()
		console.Success("All templates are up to date")
		return nil
	}

	displayUpdates(updates)
	return performUpdates(cfg, updates)
}

type updateInfo struct {
	name       string
	oldVersion string
	newVersion string
}

func checkForUpdates(cfg *config.Config, client registryv1.Client) ([]updateInfo, error) {
	var updates []updateInfo

	err := console.Spinner("Checking for updates", func() error {
		for name, installed := range cfg.Templates {
			detail, err := client.GetTemplate(name)
			if err != nil {
				console.Warning(fmt.Sprintf("Could not check updates for %s: %v", name, err))
				continue
			}

			if detail.Version != installed.Version {
				updates = append(updates, updateInfo{
					name:       name,
					oldVersion: installed.Version,
					newVersion: detail.Version,
				})
			}
		}
		return nil
	})

	return updates, err
}

func displayUpdates(updates []updateInfo) {
	console.EmptyLine()
	console.Header("Updates Available")
	console.Divider(50)
	console.EmptyLine()
	for _, u := range updates {
		console.Print("  %s %s → %s\n", console.Bold(console.Cyan(u.name)), console.Dim(u.oldVersion), console.Green(u.newVersion))
	}
	console.EmptyLine()
}

func performUpdates(cfg *config.Config, updates []updateInfo) error {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Update %d template(s)?", len(updates)),
		Default: true,
	}
	if err := survey.AskOne(prompt, &confirm); err != nil {
		return err
	}

	if !confirm {
		console.EmptyLine()
		console.Println(console.Dim("Update cancelled"))
		return nil
	}

	console.EmptyLine()
	for _, u := range updates {
		cfg.AddTemplate(u.name, u.newVersion)
		console.Print("  %s %s → %s\n", console.Green("✓"), console.Bold(u.name), console.Green(u.newVersion))
	}

	return config.Save(cfg, "")
}
