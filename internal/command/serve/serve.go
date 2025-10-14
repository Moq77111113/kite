package serve

import (
	"fmt"

	"github.com/moq77111113/kite/internal/domain/config"
	"github.com/moq77111113/kite/internal/domain/registry"
	"github.com/moq77111113/kite/internal/server"
	"github.com/moq77111113/kite/pkg/console"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	var port string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start Kite web UI and API server",
		Long:  "Starts an HTTP server that serves the Kite web UI and REST API for browsing templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load("")
			if err != nil {
				return fmt.Errorf("failed to load config: %w (run 'kite init' first)", err)
			}

			client := registry.NewClient(cfg.Registry)
			srv := server.New(client)

			console.EmptyLine()
			console.Success("Kite server starting...")
			console.EmptyLine()
			console.Print("  %s\n", console.Cyan("API Endpoints:"))
			console.Print("    http://localhost:%s/api/templates\n", port)
			console.Print("    http://localhost:%s/api/templates/:name\n", port)
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
