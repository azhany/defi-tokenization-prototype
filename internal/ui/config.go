package ui

import (
	"github.com/rivo/tview"

	"defi-tokenization-prototype/internal/config"
	"defi-tokenization-prototype/internal/eth"
)

func (t *DefiTUI) setupConfigScreen() {
	configForm := tview.NewForm()

	// Load existing config
	vault, err := config.NewVault()
	if err != nil {
		t.showError("Failed to initialize config vault: " + err.Error())
		return
	}

	cfg, err := vault.LoadConfig()
	if err != nil {
		t.showError("Failed to load config: " + err.Error())
		return
	}

	// Add form fields
	configForm.
		AddInputField("Infura Key", cfg.InfuraKey, 50, nil, nil).
		AddInputField("Keystore Path", cfg.KeystorePath, 50, nil, nil).
		AddPasswordField("Keystore Password", cfg.KeystorePass, 50, '*', nil).
		AddInputField("Token Address", cfg.TokenAddress, 50, nil, nil).
		AddInputField("Pool Address", cfg.PoolAddress, 50, nil, nil).
		AddInputField("NFT Address", cfg.NFTAddress, 50, nil, nil)

	// Add save button
	configForm.AddButton("Save", func() {
		newCfg := &config.Config{
			InfuraKey:    configForm.GetFormItemByLabel("Infura Key").(*tview.InputField).GetText(),
			KeystorePath: configForm.GetFormItemByLabel("Keystore Path").(*tview.InputField).GetText(),
			KeystorePass: configForm.GetFormItemByLabel("Keystore Password").(*tview.InputField).GetText(),
			TokenAddress: configForm.GetFormItemByLabel("Token Address").(*tview.InputField).GetText(),
			PoolAddress:  configForm.GetFormItemByLabel("Pool Address").(*tview.InputField).GetText(),
			NFTAddress:   configForm.GetFormItemByLabel("NFT Address").(*tview.InputField).GetText(),
		}

		if err := vault.SaveConfig(newCfg); err != nil {
			t.showError("Failed to save config: " + err.Error())
			return
		}

		if err := vault.SetEnvFromConfig(); err != nil {
			t.showError("Failed to set environment variables: " + err.Error())
			return
		}

		// Reinitialize Ethereum client with new config
		ethClient, err := eth.NewEthClient(
			newCfg.InfuraKey,
			newCfg.KeystorePath,
			newCfg.KeystorePass,
			newCfg.TokenAddress,
			newCfg.PoolAddress,
			newCfg.NFTAddress,
		)
		if err != nil {
			t.showError("Failed to initialize Ethereum client with new config: " + err.Error())
			return
		}

		t.ethClient = ethClient
		t.showMessage("Configuration saved successfully!")
		t.pages.SwitchToPage("menu")
	})

	// Add back button
	configForm.AddButton("Back", func() {
		t.pages.SwitchToPage("menu")
	})

	configForm.SetBorder(true).
		SetTitle("Configuration").
		SetTitleAlign(tview.AlignLeft)

	t.pages.AddPage("config", configForm, true, false)
}

// Update the showMainMenu function to include the config option
func (t *DefiTUI) showMainMenu() {
	menu := tview.NewList().
		AddItem("Configuration", "Set up your DeFi environment", 'C', func() { t.pages.SwitchToPage("config") }).
		AddItem("Mint Tokens", "Mint new stable tokens", 'M', func() { t.pages.SwitchToPage("mint") }).
		AddItem("Lending Pool", "Deposit or withdraw from lending pool", 'L', func() { t.pages.SwitchToPage("lending") }).
		AddItem("NFT Management", "Mint and manage NFTs", 'N', func() { t.pages.SwitchToPage("nft") }).
		AddItem("View Balances", "Check your token and NFT balances", 'B', func() { t.pages.SwitchToPage("balance") }).
		AddItem("Quit", "Exit the application", 'Q', t.app.Stop)

	menu.SetBorder(true).SetTitle("DeFi TUI")
	t.pages.AddPage("menu", menu, true, true)
}
