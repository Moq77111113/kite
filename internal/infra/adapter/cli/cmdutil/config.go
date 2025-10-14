package cmdutil

import (
	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/spf13/cobra"
)

// LoadConfig loads config from --config flag or default location
func LoadConfig(cmd *cobra.Command) (*config.Config, error) {
	configPath, _ := cmd.Flags().GetString("config")
	return config.Load(configPath)
}
