package cmd

import (
	"fmt"

	"github.com/moq77111113/kite/internal/config"
	"github.com/moq77111113/kite/internal/installer"

	"github.com/moq77111113/kite/internal/registry"
	"github.com/moq77111113/kite/internal/utils"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <module>",
		Short: "Add a module from the registry",
		Args:  cobra.RangeArgs(1, 10),
		RunE:  runAdd,
	}

	config, err := config.GetConfig()
	if err != nil {
		fmt.Fprintf(cmd.OutOrStderr(), "failed to get config: %v\n", err)
		return nil
	}

	cmd.Flags().StringP("registry", "r", config.Registry, "Path or URL to registry")
	cmd.Flags().StringP("out", "o", config.Out, "Output directory for the module")
	cmd.Flags().StringP("flavor", "f", config.Flavor, "Flavor of the module")
	

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) error {


	modulesNames := args

	registryPath, _ := cmd.Flags().GetString("registry")
	out, _ := cmd.Flags().GetString("out")
	flavor, _ := cmd.Flags().GetString("flavor")

	
		conf := config.Config{
			Registry: registryPath,
			Out:      out,
			Flavor:   flavor,
		}
		

	var loader registry.Loader
	if utils.IsURL(conf.Registry) {
		loader = &registry.HttpLoader{}
	} else {
		loader = &registry.FileLoader{}
	}


	reg, err := loader.LoadIndex(conf)
	if err != nil {
		return fmt.Errorf("failed to load registry index: %w", err)
	}

	fullList, err := reg.ResolveWithDependencies(args)
	if err != nil {
		return fmt.Errorf("failed to resolve module dependencies: %w", err)
	}

	modules, err := loader.LoadModules(conf, fullList)
	if err != nil {
		return fmt.Errorf("failed to load modules: %w", err)
	}

	
	if err := installer.InstallAllModules(modules, conf.Out); err != nil {
		return fmt.Errorf("failed to install modules: %w", err)
	}

	fmt.Printf("✅ Successfully installed %d module(s) into %s\n", len(modulesNames), out)
	return nil
}


