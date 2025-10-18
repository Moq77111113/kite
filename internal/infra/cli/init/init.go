package initcmd

import (
	"github.com/spf13/cobra"
)

func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project with kite.json config",
		RunE:  runInit,
	}

	cmd.Flags().StringP("path", "p", "", "Path where kits will be installed (skip interactive)")
	cmd.Flags().StringP("registry", "r", "", "Registry URL (skip interactive)")
	cmd.Flags().BoolP("force", "f", false, "Overwrite existing kite.json without prompting")

	return cmd
}

func runInit(cmd *cobra.Command, args []string) error {
	return Run(Options{
		Registry: cmd.Flag("registry").Value.String(),
		Path:     cmd.Flag("path").Value.String(),
		Force:    cmd.Flag("force").Value.String() == "true",
	})
}
