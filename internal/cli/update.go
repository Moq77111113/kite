package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/moq77111113/kite/internal/container"
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
	c, err := container.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	if len(c.Config().Templates) == 0 {
		console.EmptyLine()
		console.Println(console.Dim("No templates installed"))
		return nil
	}

	updates, err := checkForUpdates(c)
	if err != nil {
		return err
	}

	if len(updates) == 0 {
		console.EmptyLine()
		console.Success("All templates are up to date")
		return nil
	}

	displayUpdates(updates)
	return performUpdates(c, updates)
}

type updateInfo struct {
	name       string
	oldVersion string
	newVersion string
}

func checkForUpdates(c *container.Container) ([]updateInfo, error) {
	var updates []updateInfo

	err := console.Spinner("Checking for updates", func() error {
		for name, installed := range c.Config().Templates {
			detail, err := c.Client().GetTemplate(name)
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

func performUpdates(c *container.Container, updates []updateInfo) error {
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
		c.Config().AddTemplate(u.name, u.newVersion)
		console.Print("  %s %s → %s\n", console.Green("✓"), console.Bold(u.name), console.Green(u.newVersion))
	}

	return nil
}
