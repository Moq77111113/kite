package add

import (
	"fmt"

	"github.com/moq77111113/kite/internal/application/add"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

var (
	varFlags    []string
	varFileFlag string
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <kit-name> [kit-name...]",
		Short: "Download a kit from registry and add to project",
		Args:  cobra.MinimumNArgs(1),
		RunE:  runAdd,
	}
	cmd.Flags().StringArrayVarP(&varFlags, "var", "v", nil, "Set variable (key=value)")
	cmd.Flags().StringVarP(&varFileFlag, "var-file", "f", "", "Load variables from YAML file")
	return cmd
}

func runAdd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	c, ok := ctx.Value(container.ContainerKey).(*container.Container)
	if !ok || c == nil {
		return fmt.Errorf("dependencies not found in context")
	}

	providedVars, err := parseVariables(varFlags, varFileFlag)
	if err != nil {
		return fmt.Errorf("failed to parse variables: %w", err)
	}

	cfg := c.Config
	addFn := add.New(c.Installer, c.ConflictChecker, c.Tracker, c.Repository)

	console.EmptyLine()
	installed := 0
	for i, name := range args {
		if i > 0 {
			console.EmptyLine()
		}
		if err := addKit(addFn, cfg, name, providedVars); err != nil {
			console.Print("  %s %s\n", console.Red("✗"), console.Red("Failed: ", err))
			continue
		}
		installed++
	}

	if installed > 0 {
		console.EmptyLine()
		console.Success(fmt.Sprintf("Successfully installed %d kit(s)", installed))
	}
	return nil
}
