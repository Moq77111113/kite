package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/container"
	"github.com/moq77111113/kite/internal/registry"
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
	c, err := container.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	name := args[0]

	// Check if installed
	installed, ok := c.Config().GetTemplate(name)
	if !ok {
		return fmt.Errorf("template %s is not installed", name)
	}

	// Fetch from registry
	detail, err := c.Client().GetTemplate(name)
	if err != nil {
		return fmt.Errorf("failed to fetch template from registry: %w", err)
	}

	displayDiff(c, name, installed.Version, detail)
	return nil
}

func displayDiff(c *container.Container, name, localVersion string, detail *registry.TemplateDetailResponse) {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("Template: %s\n", cyan(name))
	fmt.Printf("Local version:    %s\n", localVersion)
	fmt.Printf("Registry version: %s\n", detail.Version)
	fmt.Println()

	if localVersion == detail.Version {
		fmt.Printf("%s Template is up to date\n", green("✓"))
		return
	}

	fmt.Printf("%s Update available: %s → %s\n", yellow("⚠"), localVersion, detail.Version)
	fmt.Println()

	fmt.Println("Files in registry version:")
	for _, file := range detail.Files {
		localPath := filepath.Join(c.Config().Path, name, file.Path)
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			fmt.Printf("  %s %s (new)\n", green("+"), file.Path)
		} else {
			fmt.Printf("  %s %s (modified)\n", yellow("~"), file.Path)
		}
	}

	fmt.Println()
	fmt.Println("Run 'kite update' to update this template")
}
