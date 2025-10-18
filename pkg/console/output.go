package console

import (
	"fmt"
	"strings"
	"time"
)

var portsnnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// Success prints a success message with checkmark
func Success(msg string) {
	fmt.Printf("%s %s\n", Green("✓"), Bold(msg))
}

// Error prints an error message with cross
func Error(msg string) {
	fmt.Printf("%s %s\n", Red("✗"), Red(msg))
}

// Warning prints a warning message with warning symbol
func Warning(msg string) {
	fmt.Printf("%s %s\n", Yellow("⚠"), Yellow(msg))
}

// Info prints an info message with arrow
func Info(msg string) {
	fmt.Printf("%s %s\n", Cyan("→"), Dim(msg))
}

// Header prints a bold header
func Header(text string) {
	fmt.Printf("\n%s\n", Bold(text))
}

// Divider prints a horizontal line
func Divider(length int) {
	fmt.Println(strings.Repeat("─", length))
}

// EmptyLine prints a blank line
func EmptyLine() {
	fmt.Println()
}

// Spinner runs a function with a portsnner animation
func Spinner(msg string, fn func() error) error {
	done := make(chan bool)
	errChan := make(chan error)

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r  %s %s...", Cyan(portsnnerFrames[i%len(portsnnerFrames)]), msg)
				time.Sleep(80 * time.Millisecond)
				i++
			}
		}
	}()

	go func() {
		errChan <- fn()
	}()

	err := <-errChan
	done <- true
	fmt.Print("\r\033[K")

	return err
}

// Print is a simple wrapper around fmt.Printf for consistency
func Print(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Println is a simple wrapper around fmt.Println for consistency
func Println(args ...any) {
	fmt.Println(args...)
}
