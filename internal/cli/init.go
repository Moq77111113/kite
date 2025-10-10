package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/config"
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project with kite.json config",
		RunE:  runInit,
	}

	cmd.Flags().StringP("path", "p", config.DefaultPath, "Path where templates will be installed")
	cmd.Flags().StringP("registry", "r", config.DefaultRegistry, "Registry URL")

	return cmd
}

func runInit(cmd *cobra.Command, args []string) error {

	if config.Exists() {
		return fmt.Errorf("kite.json already exists in the current directory")
	}

	path, _ := cmd.Flags().GetString("path")
	registry, _ := cmd.Flags().GetString("registry")

	_, err := config.Init(registry, path)
	if err != nil {
		return fmt.Errorf("failed to initialize config: %w", err)
	}

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s Created kite.json\n", green("✓"))
	fmt.Printf("%s Created %s directory\n", green("✓"), path)
	fmt.Printf("%s Ready to install templates! Run 'kite list' to see available templates.\n", green("✓"))

	return nil
}
