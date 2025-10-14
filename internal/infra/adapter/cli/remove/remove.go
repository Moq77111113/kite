package remove

import (
	"fmt"

	"github.com/moq77111113/kite/internal/infra/adapter/cli/cmdutil"
	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewRemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <template-name>",
		Short: "Remove an installed template",
		Args:  cobra.ExactArgs(1),
		RunE:  runRemove,
	}
}

func runRemove(cmd *cobra.Command, args []string) error {
	cfg, err := cmdutil.LoadConfig(cmd)
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	svc := template.NewService(cfg)

	name := args[0]
	if err := svc.Remove(name); err != nil {
		return err
	}

	console.EmptyLine()
	console.Success(fmt.Sprintf("Removed %s", name))

	return nil
}
