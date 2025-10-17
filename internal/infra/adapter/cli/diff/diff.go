package diff

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/moq77111113/kite/internal/infra/adapter/cli/cmdutil"
	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/infra/persistence/config"
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
	cfg, err := cmdutil.LoadConfig(cmd)
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	name := args[0]

	svc := template.NewService(cfg, nil)

	installed, err := svc.GetInstalled(name)
	if err != nil {
		return err
	}

	detail, err := svc.GetDetails(name)
	if err != nil {
		return fmt.Errorf("failed to fetch template from registry: %w", err)
	}

	displayDiff(cfg, name, installed.Version, detail)
	return nil
}

func displayDiff(cfg *config.Config, name, localVersion string, detail *registry.TemplateDetailResponse) {
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
