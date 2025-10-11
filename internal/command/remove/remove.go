package remove

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/config"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/domain/template"
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
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	client := registry.NewClient(cfg.Registry)
	mgr := template.NewManager(cfg, client)

	name := args[0]
	if err := mgr.Remove(name); err != nil {
		return err
	}

	console.EmptyLine()
	console.Success(fmt.Sprintf("Removed %s", name))

	return nil
}
