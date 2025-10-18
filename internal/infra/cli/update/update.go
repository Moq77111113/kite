package update

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	updateapp "github.com/moq77111113/kite/internal/application/update"

	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update installed kits to latest versions",
		RunE:  runUpdate,
	}
}

func runUpdate(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	c, ok := ctx.Value(container.ContainerKey).(*container.Container)
	if !ok || c == nil {
		return fmt.Errorf("dependencies not found in context")
	}
	cfg := c.Config
	if len(cfg.Kits) == 0 {
		console.EmptyLine()
		console.Println(console.Dim("No kits installed"))
		return nil
	}

	updateSvc := updateapp.New(
		c.Repository,
		c.InstallationRegistry,
		c.VersionComparator,
		c.KitLifecycle,
	)
	var updates []updateapp.Check
	err := console.Spinner("Checking for updates", func() error {
		var checkErr error
		updates, checkErr = updateSvc.CheckAll()
		return checkErr
	})
	if err != nil {
		return err
	}
	if len(updates) == 0 {
		console.EmptyLine()
		console.Success("All kits are up to date")
		return nil
	}
	displayUpdates(updates)
	if !promptConfirmation(len(updates)) {
		console.EmptyLine()
		console.Println(console.Dim("Update cancelled"))
		return nil
	}
	console.EmptyLine()
	for _, u := range updates {
		err := updateSvc.ApplyUpdate(u.Name, cfg.Path)
		if err != nil {
			console.Print("  %s %s: %v\n", console.Red("✗"), console.Bold(u.Name), err)
			continue
		}
		console.Print("  %s %s %s → %s\n",
			console.Green("✓"),
			console.Bold(u.Name),
			console.Dim(u.CurrentVersion),
			console.Green(u.LatestVersion),
		)
	}
	return nil
}

func displayUpdates(updates []updateapp.Check) {
	console.EmptyLine()
	console.Header("Updates Available")
	console.Divider(50)
	console.EmptyLine()
	for _, u := range updates {
		console.Print("  %s %s → %s\n",
			console.Bold(console.Cyan(u.Name)),
			console.Dim(u.CurrentVersion),
			console.Green(u.LatestVersion))
	}
	console.EmptyLine()
}

func promptConfirmation(count int) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Update %d kit(s)?", count),
		Default: true,
	}
	if err := survey.AskOne(prompt, &confirm); err != nil {
		return false
	}
	return confirm
}
