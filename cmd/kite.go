/*
Copyright © 2025 Moq
*/
package main

import (
	"os"

	"github.com/moq77111113/kite/pkg/cmd"
)



func main() {
	root := cmd.NewKiteCmd()

	root.AddCommand(cmd.NewAddCmd())


	if err := root.Execute(); err != nil {
		os.Exit(1)
	}


}


