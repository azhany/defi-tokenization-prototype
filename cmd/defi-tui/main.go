package main

import (
	"log"

	"defi-tokenization-prototype/internal/config"
	"defi-tokenization-prototype/internal/ui"
)

func main() {
	// Initialize vault and load config
	vault, err := config.NewVault()
	if err != nil {
		log.Fatal("Failed to initialize config vault:", err)
	}

	// Load and set environment variables from config
	if err := vault.SetEnvFromConfig(); err != nil {
		log.Printf("No existing configuration found, you'll need to set it up in the UI")
	}

	// Create TUI without Ethereum client initially
	tui := ui.NewDefiTUI(nil)

	// The Ethereum client will be initialized after configuration is set in the UI
	if err := tui.Run(); err != nil {
		log.Fatal("Error running TUI:", err)
	}
}
