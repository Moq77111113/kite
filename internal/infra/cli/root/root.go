package root

import (
	"context"

	"github.com/moq77111113/kite/internal/infra/cli/add"
	"github.com/moq77111113/kite/internal/infra/cli/annotations"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	initcmd "github.com/moq77111113/kite/internal/infra/cli/init"
	"github.com/moq77111113/kite/internal/infra/cli/list"
	"github.com/moq77111113/kite/internal/infra/cli/remove"
	"github.com/moq77111113/kite/internal/infra/cli/serve"
	"github.com/moq77111113/kite/internal/infra/cli/update"
	"github.com/moq77111113/kite/internal/infra/cli/version"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kite",
		Short: "Infrastructure kit manager",
		Long:  `Kite is a CLI tool for managing infrastructure kits. Pull kits from a registry.`,
	}

	cmd.PersistentFlags().StringP("config", "c", "", "Config file path (default: ./kite.yaml or ~/.kite/config.yaml)")

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if cmd.Annotations[annotations.SkipContainer.String()] == "true" {
			return nil
		}

		cfgPath, _ := cmd.Flags().GetString("config")
		c, err := container.NewContainer(cfgPath)
		if err != nil {
			return err
		}
		ctx := context.WithValue(cmd.Context(), container.ContainerKey, c)
		cmd.SetContext(ctx)
		return nil
	}

	cmd.AddCommand(initcmd.NewInitCmd())
	cmd.AddCommand(add.NewAddCmd())
	cmd.AddCommand(list.NewListCmd())
	cmd.AddCommand(remove.NewRemoveCmd())
	cmd.AddCommand(update.NewUpdateCmd())
	cmd.AddCommand(serve.NewServeCmd())
	cmd.AddCommand(version.NewVersionCmd())

	return cmd
}
