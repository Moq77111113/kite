package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/moq77111113/kite/internal/container"
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
	c, err := container.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
	}

	name := args[0]
	if err := c.Manager().Remove(name); err != nil {
		return err
	}

	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s Removed %s\n", green("✓"), name)
	fmt.Printf("%s Updated kite.json\n", green("✓"))

	return nil
}
