package diff

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/domain/config"
	"github.com/moq77111113/kite/internal/domain/registry"
registryv1 "github.com/moq77111113/kite/api/registry/v1"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewDiffCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "diff <template-name>",
		Short: "Show what changed between local and registry version",
		Args:  cobra.ExactArgs(1),
		RunE:  runDiff,
	}
}

func runDiff(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	name := args[0]

	// Check if installed
	installed, ok := cfg.GetTemplate(name)
	if !ok {
		return fmt.Errorf("template %s is not installed", name)
	}

	// Fetch from registry
	client := registry.NewClient(cfg.Registry)
	detail, err := client.GetTemplate(name)
	if err != nil {
		return fmt.Errorf("failed to fetch template from registry: %w", err)
	}

	displayDiff(cfg, name, installed.Version, detail)
	return nil
}

func displayDiff(cfg *config.Config, name, localVersion string, detail *registryv1.TemplateDetailResponse) {
	console.EmptyLine()
	console.Print("%s %s\n", console.Bold("Template:"), console.Cyan(name))
	console.Divider(50)
	console.EmptyLine()
	console.Print("  Local:    %s\n", console.Dim(localVersion))
	console.Print("  Registry: %s\n", console.Cyan(detail.Version))
	console.EmptyLine()

	if localVersion == detail.Version {
		console.Success("Template is up to date")
		return
	}

	console.Warning(fmt.Sprintf("Update available: %s → %s", localVersion, detail.Version))
	console.EmptyLine()
	console.Header("Changes:")
	console.EmptyLine()

	for _, file := range detail.Files {
		localPath := filepath.Join(cfg.Path, name, file.Path)
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			console.Print("  %s %s %s\n", console.Green("+"), file.Path, console.Dim("(new)"))
		} else {
			console.Print("  %s %s %s\n", console.Yellow("~"), file.Path, console.Dim("(modified)"))
		}
	}

	console.EmptyLine()
	console.Divider(50)
	console.Info("Run 'kite update' to apply changes")
	console.EmptyLine()
}
