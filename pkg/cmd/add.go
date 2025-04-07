package cmd

import (
	"fmt"
	"strings"

	"github.com/moq77111113/kite/internal/installer"
	"github.com/moq77111113/kite/internal/registry"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <module>",
		Short: "Add a module from the registry",
		Args:  cobra.ExactArgs(1),
		RunE:  runAdd,
	}

	cmd.Flags().StringP("registry", "r", "./registry.json", "Path or URL to registry")
	cmd.Flags().StringP("out", "o", "./", "Destination directory for the module")

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) error {
	moduleName := args[0]

	registryPath, _ := cmd.Flags().GetString("registry")
	destDir, _ := cmd.Flags().GetString("out")

	var loader registry.Loader
	if isURL(registryPath) {
		loader = &registry.HttpLoader{}
	} else {
		loader = &registry.FileLoader{}
	}

	reg, err := loader.Load(registryPath)
	if err != nil {
		return fmt.Errorf("failed to load registry: %w", err)
	}

	module := reg.FindByName(moduleName)
	if module == nil {
		return fmt.Errorf("module %q not found in registry", moduleName)
	}

	if err := installer.InstallModule(module, destDir); err != nil {
		return fmt.Errorf("failed to install module: %w", err)
	}

	fmt.Printf("✅ Module %q installed to %s\n", moduleName, destDir)
	return nil
}

func isURL(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}
