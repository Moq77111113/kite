package cmd

import (
	"fmt"
	"strings"

	"github.com/moq77111113/kite/internal/config"
	"github.com/moq77111113/kite/internal/installer"
	"github.com/moq77111113/kite/internal/registry"
	"github.com/moq77111113/kite/internal/utils"
	spinner "github.com/moq77111113/kite/internal/vendors"
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
	registryPath, _ := cmd.Flags().GetString("registry")
	out, _ := cmd.Flags().GetString("out")
	flavor, _ := cmd.Flags().GetString("flavor")

	conf := config.Config{
		Registry: registryPath,
		Out:      out,
		Flavor:   flavor,
	}

	var reg *registry.Registry
	var fullList []string
	var modules []*registry.Module


	steps := []spinner.Step{
		{
			Name: "Initializing loader",
			Action: func() error {
				return nil
			},
		},
		{
			Name: "Loading registry index",
			Action: func() error {
				var err error
				reg, err = loadRegistryIndex(conf)
				return err
			},
		},
		{
			Name: "Resolving dependencies",
			Action: func() error {
				var err error
				fullList, err = resolveDependencies(reg, args)
				return err
			},
		},
		{
			Name: "Loading modules",
			Action: func() error {
				var err error
				modules, err = loadModules(conf, fullList)
				return err
			},
		},
		{
			Name: "Installing modules",
			Action: func() error {
				return installModules(modules, conf.Out)
			},
		},
	}

	if err := spinner.RunSteps(steps, "Adding modules"); err != nil {
		return err
	}

	fmt.Printf("✅ Successfully installed %d module(s): %s into %s\n", len(fullList), strings.Join(fullList, ", "), out)
	return nil
}

func loadRegistryIndex(conf config.Config) (*registry.Registry, error) {
	var loader registry.Loader
	if utils.IsURL(conf.Registry) {
		loader = &registry.HttpLoader{}
	} else {
		loader = &registry.FileLoader{}
	}
	
	return loader.LoadIndex(conf)
}

func resolveDependencies(reg *registry.Registry, modules []string) ([]string, error) {
	return reg.ResolveWithDependencies(modules)
}

func loadModules(conf config.Config, moduleNames []string) ([]*registry.Module, error) {
	var loader registry.Loader
	if utils.IsURL(conf.Registry) {
		loader = &registry.HttpLoader{}
	} else {
		loader = &registry.FileLoader{}
	}
	
	return loader.LoadModules(conf, moduleNames)
}

func installModules(modules []*registry.Module, destDir string) error {
	return installer.InstallAllModules(modules, destDir)
}