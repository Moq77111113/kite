package cmd

import (
	"fmt"
	"os"

	"github.com/moq77111113/kite/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewKiteCmd() *cobra.Command {
	cobra.OnInitialize(loadConfig)
	return &cobra.Command{
		Use:   "kite",
		Short: "Friction free, flat module sharing",
		Long: `Kite is a minimalist CLI for importing flat files - 
	config, CI, snippets, tooling - from a centralized registry. 
	No more manual copy and paste, kite takes care of everything.`,
	}
}


func loadConfig() {
	err := config.Load()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return 
		}
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}
}