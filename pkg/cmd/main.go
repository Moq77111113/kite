package cmd


import (
	"github.com/spf13/cobra"
)

func NewKiteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "kite",
		Short: "Friction free, flat module sharing",
		Long: `Kite is a minimalist CLI for importing flat files - 
	config, CI, snippets, tooling - from a centralized registry. 
	No more manual copy and paste, kite takes care of everything.`,
	}
}