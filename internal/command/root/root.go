package root

import (
	"github.com/moq77111113/kite/internal/command/add"
	"github.com/moq77111113/kite/internal/command/diff"
	initcmd "github.com/moq77111113/kite/internal/command/init"
	"github.com/moq77111113/kite/internal/command/list"
	"github.com/moq77111113/kite/internal/command/remove"
	"github.com/moq77111113/kite/internal/command/serve"
	"github.com/moq77111113/kite/internal/command/update"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kite",
		Short: "Infrastructure template manager",
		Long: `Kite is a CLI tool for managing infrastructure templates.
Pull templates from a registry.`,
	}

	cmd.PersistentFlags().StringP("config", "c", "", "Config file path (default: ./kite.yaml or ~/.kite/config.yaml)")

	cmd.AddCommand(initcmd.NewInitCmd())
	cmd.AddCommand(add.NewAddCmd())
	cmd.AddCommand(list.NewListCmd())
	cmd.AddCommand(remove.NewRemoveCmd())
	cmd.AddCommand(update.NewUpdateCmd())
	cmd.AddCommand(diff.NewDiffCmd())
	cmd.AddCommand(serve.NewServeCmd())

	return cmd
}
