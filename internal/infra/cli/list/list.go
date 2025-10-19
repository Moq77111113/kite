package list

import (
	"fmt"
	"strings"

	listapp "github.com/moq77111113/kite/internal/application/list"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available kits from registry",
		RunE:  runList,
	}
}

func runList(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	c, ok := ctx.Value(container.ContainerKey).(*container.Container)
	if !ok || c == nil {
		return fmt.Errorf("dependencies not found in context")
	}

	ls := listapp.New(c.Repository, c.Tracker)

	kits, err := ls.Execute()
	if err != nil {
		return fmt.Errorf("failed to fetch kits: %w", err)
	}

	displayKits(kits)
	return nil
}

func displayKits(kits []listapp.Item) {
	console.EmptyLine()
	console.Header("Available Kits")
	console.Divider(50)
	console.EmptyLine()

	for _, t := range kits {
		statusIcon := console.Dim("○")
		if t.Installed {
			statusIcon = console.Green("●")
		}

		console.Print("  %s %s %s\n",
			statusIcon,
			console.Bold(console.Cyan(t.Name)),
			console.Dim(fmt.Sprintf("v%s", t.Version)))
		console.Print("     %s\n", t.Description)
		if len(t.Tags) > 0 {
			console.Print("     %s\n", console.Gray(strings.Join(t.Tags, " · ")))
		}
		console.EmptyLine()
	}

	console.Divider(50)
	console.Info("Run 'kite add <kit>' to install")
	console.EmptyLine()
}
