package main

import (
	"os"

	"github.com/moq77111113/kite/internal/cli"
)

func main() {
	root := cli.NewRootCmd()

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
