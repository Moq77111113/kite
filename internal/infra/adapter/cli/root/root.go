package root

import (
	"github.com/moq77111113/kite/internal/infra/adapter/cli/add"
	"github.com/moq77111113/kite/internal/infra/adapter/cli/diff"
	initcmd "github.com/moq77111113/kite/internal/infra/adapter/cli/init"
	"github.com/moq77111113/kite/internal/infra/adapter/cli/list"
	"github.com/moq77111113/kite/internal/infra/adapter/cli/remove"
	"github.com/moq77111113/kite/internal/infra/adapter/cli/serve"
	"github.com/moq77111113/kite/internal/infra/adapter/cli/update"
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
