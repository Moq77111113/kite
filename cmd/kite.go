package main

import (
	"os"

	"github.com/moq77111113/kite/internal/command/root"
)

func main() {
	rootCmd := root.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
