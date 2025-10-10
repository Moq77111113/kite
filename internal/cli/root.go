package cli

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kite",
		Short: "Infrastructure template manager",
		Long: `Kite is a CLI tool for managing infrastructure templates.
Pull templates from a registry.`,
	}

	cmd.AddCommand(NewInitCmd())
	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewRemoveCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewDiffCmd())

	return cmd
}
