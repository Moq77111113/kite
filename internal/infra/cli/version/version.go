package version

import (
	"github.com/moq77111113/kite/internal/version"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		RunE:  runVersion,
	}
}

func runVersion(cmd *cobra.Command, args []string) error {
	console.EmptyLine()
	console.Header("Kite Version")
	console.Divider(50)
	console.EmptyLine()

	console.Print("  %s: %s\n", console.Bold("Version"), console.Cyan(version.GetVersion()))
	console.Print("  %s: %s\n", console.Bold("Commit"), console.Dim(version.GetCommit()))
	console.Print("  %s: %s\n", console.Bold("Built"), console.Dim(version.GetDate()))

	console.EmptyLine()
	console.Divider(50)
	console.EmptyLine()

	return nil
}
