package cli

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/container"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add <template-name> [template-name...]",
		Short: "Download a template from registry and add to project",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runAdd,
	}
}

func runAdd(cmd *cobra.Command, args []string) error {
	c, err := container.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	installed := 0
	for _, name := range args {
		if err := addTemplate(c, name); err != nil {
			if err.Error() == fmt.Sprintf("template %s is already installed", name) {
				fmt.Printf("%s %s is already installed\n", yellow("⚠"), name)
			} else {
				color.Red("✗ Failed to add %s: %v", name, err)
			}
			continue
		}
		installed++
	}

	if installed > 0 {
		fmt.Printf("\n%s Successfully installed %d template(s)\n", green("✓"), installed)
	}

	return nil
}

func addTemplate(c *container.Container, name string) error {
	fmt.Printf("⠿ Fetching %s from registry...\n", name)

	if err := c.Manager().Add(name); err != nil {
		return err
	}

	green := color.New(color.FgGreen).SprintFunc()
	destPath := filepath.Join(c.Config().Path, name)
	fmt.Printf("%s Installed to %s\n", green("✓"), destPath)
	fmt.Printf("%s Updated kite.json\n", green("✓"))

	return nil
}
