package list

import (
	"fmt"
	"strings"

	"github.com/moq77111113/kite/internal/infra/adapter/cli/cmdutil"
	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
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
	cfg, err := cmdutil.LoadConfig(cmd)
	if err != nil {
		return err
	}

	svc := template.NewService(cfg, nil)
	templates, err := svc.ListAvailable()
	if err != nil {
		return fmt.Errorf("failed to fetch templates: %w", err)
	}

	displayTemplates(cfg, templates)
	return nil
}

func displayTemplates(cfg *config.Config, templates []registry.TemplateSummary) {
	console.EmptyLine()
	console.Header("Available Templates")
	console.Divider(50)
	console.EmptyLine()

	for _, t := range templates {
		_, installed := cfg.GetTemplate(t.Name)

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
