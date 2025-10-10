package cli

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/container"
	"github.com/moq77111113/kite/internal/registry"
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
	fmt.Println("Available templates:")
	fmt.Println()

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	gray := color.New(color.FgHiBlack).SprintFunc()

	for _, t := range templates {
		_, installed := c.Config().GetTemplate(t.Name)
		status := red("✗ Not installed")
		if installed {
			status = green("✓ Installed")
		}

		fmt.Printf("  %s (%s) - %s\n", cyan(t.Name), t.Version, t.Description)
		if len(t.Tags) > 0 {
			fmt.Printf("    Tags: %s\n", gray(strings.Join(t.Tags, ", ")))
		}
		fmt.Printf("    %s\n", status)
		fmt.Println()
	}

	fmt.Println("Run 'kite add <template>' to install")
}
