package cli

import (
	"github.com/spf13/cobra"

	"github.com/moq77111113/kite/internal/initialize"
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project with kite.json config",
		RunE:  runInit,
	}

	cmd.Flags().StringP("path", "p", "", "Path where templates will be installed (skip interactive)")
	cmd.Flags().StringP("registry", "r", "", "Registry URL (skip interactive)")
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing kite.json without prompting")

	return cmd
}

func runInit(cmd *cobra.Command, args []string) error {
	return initialize.Run(initialize.Options{
		Registry: cmd.Flag("registry").Value.String(),
		Path:     cmd.Flag("path").Value.String(),
		Force:    cmd.Flag("force").Value.String() == "true",
	})
}
