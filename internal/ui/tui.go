package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"defi-tokenization-prototype/internal/eth"
)

type DefiTUI struct {
	app          *tview.Application
	pages        *tview.Pages
	mintForm     *tview.Form
	lendForm     *tview.Form
	nftForm      *tview.Form
	balanceTable *tview.Table
	ethClient    *eth.EthClient
}

func NewDefiTUI(ethClient *eth.EthClient) *DefiTUI {
	tui := &DefiTUI{
		app:       tview.NewApplication(),
		pages:     tview.NewPages(),
		ethClient: ethClient,
	}

	tui.setupMintScreen()
	tui.setupLendingScreen()
	tui.setupNFTScreen()
	tui.setupBalanceScreen()

	return tui
}

func (t *DefiTUI) setupMintScreen() {
	t.mintForm = tview.NewForm().
		AddInputField("Amount", "", 20, nil, nil).
		AddButton("Mint", t.handleMint).
		AddButton("Back", t.showMainMenu)

	t.mintForm.SetBorder(true).SetTitle("Mint Stable Token")
	t.pages.AddPage("mint", t.mintForm, true, false)
}

func (t *DefiTUI) setupLendingScreen() {
	t.lendForm = tview.NewForm().
		AddInputField("Amount", "", 20, nil, nil).
		AddButton("Deposit", t.handleDeposit).
		AddButton("Withdraw", t.handleWithdraw).
		AddButton("Back", t.showMainMenu)

	t.lendForm.SetBorder(true).SetTitle("Lending Pool")
	t.pages.AddPage("lending", t.lendForm, true, false)
}

func (t *DefiTUI) setupNFTScreen() {
	t.nftForm = tview.NewForm().
		AddInputField("Collateral Amount", "", 20, nil, nil).
		AddButton("Mint NFT", t.handleNFTMint).
		AddButton("Back", t.showMainMenu)

	t.nftForm.SetBorder(true).SetTitle("NFT Management")
	t.pages.AddPage("nft", t.nftForm, true, false)
}

func (t *DefiTUI) setupBalanceScreen() {
	t.balanceTable = tview.NewTable().
		SetBorders(true)

	t.balanceTable.SetCell(0, 0, tview.NewTableCell("Asset").SetTextColor(tcell.ColorYellow))
	t.balanceTable.SetCell(0, 1, tview.NewTableCell("Balance").SetTextColor(tcell.ColorYellow))

	t.pages.AddPage("balance", t.balanceTable, true, false)
}

func (t *DefiTUI) showMainMenu() {
	menu := tview.NewList().
		AddItem("Mint Tokens", "Mint new stable tokens", 'M', func() { t.pages.SwitchToPage("mint") }).
		AddItem("Lending Pool", "Deposit or withdraw from lending pool", 'L', func() { t.pages.SwitchToPage("lending") }).
		AddItem("NFT Management", "Mint and manage NFTs", 'N', func() { t.pages.SwitchToPage("nft") }).
		AddItem("View Balances", "Check your token and NFT balances", 'B', func() { t.pages.SwitchToPage("balance") }).
		AddItem("Quit", "Exit the application", 'Q', t.app.Stop)

	menu.SetBorder(true).SetTitle("DeFi TUI")
	t.pages.AddPage("menu", menu, true, true)
}

func (t *DefiTUI) Run() error {
	t.showMainMenu()
	return t.app.SetRoot(t.pages, true).Run()
}
