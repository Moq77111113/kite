package console

import "github.com/fatih/color"

// Color functions for consistent styling across the application
var (
	Green  = color.New(color.FgGreen).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	Bold   = color.New(color.Bold).SprintFunc()
	Dim    = color.New(color.Faint).SprintFunc()
	Gray   = color.New(color.FgHiBlack).SprintFunc()
)
