package cli

import (
	"fmt"
	"strings"

	"github.com/moq77111113/kite/internal/container"
	"github.com/moq77111113/kite/internal/registry"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all available templates from registry",
		RunE:  runList,
	}
}

func runList(cmd *cobra.Command, args []string) error {
	c, err := container.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	templates, err := c.Client().ListTemplates()
	if err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	displayTemplates(c, templates)
	return nil
}

func displayTemplates(c *container.Container, templates []registry.TemplateSummary) {
	console.EmptyLine()
	console.Header("Available Templates")
	console.Divider(50)
	console.EmptyLine()

	for _, t := range templates {
		_, installed := c.Config().GetTemplate(t.Name)

		statusIcon := console.Dim("○")
		if installed {
			statusIcon = console.Green("●")
		}

		console.Print("  %s %s %s\n", statusIcon, console.Bold(console.Cyan(t.Name)), console.Dim(fmt.Sprintf("v%s", t.Version)))
		console.Print("     %s\n", t.Description)
		if len(t.Tags) > 0 {
			console.Print("     %s\n", console.Gray(strings.Join(t.Tags, " · ")))
		}
		console.EmptyLine()
	}

	console.Divider(50)
	console.Info("Run 'kite add <template>' to install")
	console.EmptyLine()
}
