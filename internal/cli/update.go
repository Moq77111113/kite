package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/container"
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
		fmt.Println("No templates installed")
		return nil
	}

	updates, err := checkForUpdates(c)
	if err != nil {
		return err
	}

	if len(updates) == 0 {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s All templates are up to date\n", green("✓"))
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
	fmt.Println("Checking for updates...")
	fmt.Println()

	var updates []updateInfo

	for name, installed := range c.Config().Templates {
		detail, err := c.Client().GetTemplate(name)
		if err != nil {
			color.Yellow("⚠ Could not check updates for %s: %v", name, err)
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

	return updates, nil
}

func displayUpdates(updates []updateInfo) {
	fmt.Println("Updates available:")
	for _, u := range updates {
		fmt.Printf("  %s: %s → %s\n", u.name, u.oldVersion, u.newVersion)
	}
	fmt.Println()
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
		fmt.Println("Update cancelled")
		return nil
	}

	green := color.New(color.FgGreen).SprintFunc()
	for _, u := range updates {
		c.Config().AddTemplate(u.name, u.newVersion)
		fmt.Printf("%s Updated %s to %s\n", green("✓"), u.name, u.newVersion)
	}

	return nil
}
