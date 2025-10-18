package serve

import (
	"fmt"

	"github.com/moq77111113/kite/internal/infra/api"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	var port string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start Kite web UI and API server",
		Long:  "Starts an HTTP server that serves the Kite web UI and REST API for browsing kits",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			c, ok := ctx.Value(container.ContainerKey).(*container.Container)
			if !ok || c == nil {
				return fmt.Errorf("dependencies not found in context")
			}

			srv := api.New(c)

			console.EmptyLine()
			console.Success("Kite server starting...")
			console.EmptyLine()
			console.Print("  %s\n", console.Cyan("Web UI:"))
			console.Print("    http://localhost:%s/\n", port)
			console.EmptyLine()
			console.Print("  %s\n", console.Cyan("API Endpoints:"))
			console.Print("    http://localhost:%s/api/kits\n", port)
			console.Print("    http://localhost:%s/api/kits/:name\n", port)
			console.EmptyLine()
			console.Print("  %s\n", console.Yellow("Press Ctrl+C to stop"))
			console.EmptyLine()

			if err := srv.Start(port); err != nil {
				return fmt.Errorf("server error: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen on")

	return cmd
}
