package initcmd

import (
	"fmt"

	"github.com/moq77111113/kite/internal/infra/persistence/config"
	"github.com/moq77111113/kite/pkg/console"
)

func showWelcome() {
	console.EmptyLine()
	console.Println(console.Bold("🪁 Welcome to Kite!"))
	console.EmptyLine()
	console.Println("Let's set up your project for infrastructure kits.")
	console.EmptyLine()
}

func promptUpdate(cfg *config.Config) (bool, error) {
	configPath := config.GetConfigPath("")
	console.EmptyLine()
	console.Warning(fmt.Sprintf("Config already exists at %s", configPath))
	console.EmptyLine()
	console.Print("  Registry: %s\n", console.Cyan(cfg.Registry))
	console.Print("  Path:     %s\n", console.Cyan(cfg.Path))
	if len(cfg.Kits) > 0 {
		console.Print("  Kits: %s\n", console.Dim(fmt.Sprintf("%d installed", len(cfg.Kits))))
	}
	console.EmptyLine()

	return AskConfirm("Do you want to update the configuration?", false)
}

func showUpdateSuccess(cfg *config.Config, registry, path string) {
	configPath := config.GetConfigPath("")
	console.EmptyLine()
	console.Success(fmt.Sprintf("Updated config at %s", configPath))
	console.Print("%s %s\n", console.Green("✓"), fmt.Sprintf("Registry: %s", console.Cyan(registry)))
	console.Print("%s %s\n", console.Green("✓"), fmt.Sprintf("Path: %s", console.Cyan(path)))
	if len(cfg.Kits) > 0 {
		console.Print("%s %s\n", console.Green("✓"), fmt.Sprintf("Preserved %d installed kit(s)", len(cfg.Kits)))
	}
	console.EmptyLine()
	console.Info("Run 'kite list' to see available kits")
}

func showCreateSuccess(registry, path string) {
	configPath := config.GetConfigPath("")
	console.EmptyLine()
	console.Success(fmt.Sprintf("Created config at %s", configPath))
	console.Print("%s %s\n", console.Green("✓"), fmt.Sprintf("Created %s directory", console.Cyan(path)))
	console.Print("%s %s\n", console.Green("✓"), fmt.Sprintf("Using registry: %s", console.Cyan(registry)))
	console.EmptyLine()
	console.Info("Run 'kite list' to see available kits")
}
