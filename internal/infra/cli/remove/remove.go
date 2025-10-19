package remove

import (
	"fmt"

	"github.com/moq77111113/kite/internal/application/remove"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <kit-name>",
		Short: "Remove an installed kit",
		Args:  cobra.ExactArgs(1),
		RunE:  runRemove,
	}
}

func runRemove(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	c, ok := ctx.Value(container.ContainerKey).(*container.Container)
	if !ok || c == nil {
		return fmt.Errorf("dependencies not found in context")
	}

	cfg := c.Config

	rm := remove.New(c.Manager)

	name := args[0]
	if err := rm.Execute(name, cfg.Path); err != nil {
		return err
	}

	console.EmptyLine()
	console.Success(fmt.Sprintf("Removed %s", name))

	return nil
}
