package ui

import (
	"math/big"
	"strconv"

	"defi-tokenization-prototype/internal/eth"
)

type Handlers struct {
	ethClient *eth.EthClient
	tui       *DefiTUI
}

func (t *DefiTUI) handleMint() {
	amount := t.mintForm.GetFormItem(0).(*tview.InputField).GetText()
	value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		t.showError("Invalid amount")
		return
	}

	go func() {
		_, err := t.ethClient.StableToken.Mint(t.ethClient.Auth, big.NewInt(value))
		if err != nil {
			t.showError("Failed to mint tokens: " + err.Error())
			return
		}
		t.showMessage("Successfully minted " + amount + " tokens")
	}()
}

func (t *DefiTUI) handleDeposit() {
	amount := t.lendForm.GetFormItem(0).(*tview.InputField).GetText()
	value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		t.showError("Invalid amount")
		return
	}

	go func() {
		_, err := t.ethClient.LendingPool.Deposit(t.ethClient.Auth, big.NewInt(value))
		if err != nil {
			t.showError("Failed to deposit: " + err.Error())
			return
		}
		t.showMessage("Successfully deposited " + amount + " tokens")
	}()
}

func (t *DefiTUI) handleWithdraw() {
	amount := t.lendForm.GetFormItem(0).(*tview.InputField).GetText()
	value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		t.showError("Invalid amount")
		return
	}

	go func() {
		_, err := t.ethClient.LendingPool.Withdraw(t.ethClient.Auth, big.NewInt(value))
		if err != nil {
			t.showError("Failed to withdraw: " + err.Error())
			return
		}
		t.showMessage("Successfully withdrawn " + amount + " tokens")
	}()
}

func (t *DefiTUI) handleNFTMint() {
	amount := t.nftForm.GetFormItem(0).(*tview.InputField).GetText()
	value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		t.showError("Invalid amount")
		return
	}

	go func() {
		_, err := t.ethClient.NFTContract.Mint(t.ethClient.Auth, big.NewInt(value))
		if err != nil {
			t.showError("Failed to mint NFT: " + err.Error())
			return
		}
		t.showMessage("Successfully minted NFT with " + amount + " collateral")
	}()
}

func (t *DefiTUI) showError(msg string) {
	t.app.QueueUpdateDraw(func() {
		modal := tview.NewModal().
			SetText(msg).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				t.pages.RemovePage("error")
			})
		t.pages.AddPage("error", modal, true, true)
	})
}

func (t *DefiTUI) showMessage(msg string) {
	t.app.QueueUpdateDraw(func() {
		modal := tview.NewModal().
			SetText(msg).
			AddButtons([]string{"OK"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				t.pages.RemovePage("message")
			})
		t.pages.AddPage("message", modal, true, true)
	})
}

func (t *DefiTUI) updateBalances() {
	go func() {
		tokenBalance, err := t.ethClient.StableToken.BalanceOf(nil, t.ethClient.Auth.From)
		if err != nil {
			t.showError("Failed to fetch token balance")
			return
		}

		poolBalance, err := t.ethClient.LendingPool.GetUserBalance(nil, t.ethClient.Auth.From)
		if err != nil {
			t.showError("Failed to fetch pool balance")
			return
		}

		nfts, err := t.ethClient.NFTContract.TokensOfOwner(nil, t.ethClient.Auth.From)
		if err != nil {
			t.showError("Failed to fetch NFTs")
			return
		}

		t.app.QueueUpdateDraw(func() {
			t.balanceTable.SetCell(1, 0, tview.NewTableCell("Stable Token"))
			t.balanceTable.SetCell(1, 1, tview.NewTableCell(tokenBalance.String()))
			t.balanceTable.SetCell(2, 0, tview.NewTableCell("Lending Pool"))
			t.balanceTable.SetCell(2, 1, tview.NewTableCell(poolBalance.String()))
			t.balanceTable.SetCell(3, 0, tview.NewTableCell("NFTs Owned"))
			t.balanceTable.SetCell(3, 1, tview.NewTableCell(strconv.Itoa(len(nfts))))
		})
	}()
}
