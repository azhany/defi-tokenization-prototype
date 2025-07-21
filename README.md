# DeFi Tokenization TUI

A terminal-based user interface for interacting with DeFi smart contracts on the Ethereum testnet.

## Features

- Mint stable tokens
- Deposit/withdraw from lending pool
- Mint NFTs with collateral
- View token balances and NFT positions

## Prerequisites

- Go 1.21 or later
- An Ethereum testnet account (Goerli)
- Access to deployed smart contracts (token, lending pool, and NFT)

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

## Configuration

### Interactive Configuration & Secure Vault

On first run, or anytime via the "Configuration" menu, you can set all required settings interactively:

- **Infura Key**: Your Infura project key
- **Keystore Path**: Path to your Ethereum keystore file
- **Keystore Password**: Password for your keystore
- **Token Address**: Deployed stable token contract address
- **Pool Address**: Deployed lending pool contract address
- **NFT Address**: Deployed NFT contract address

All sensitive data is stored securely using your system's keyring (Windows Credential Manager, macOS Keychain, Linux Secret Service) or encrypted file vault if unavailable. No need to set environment variables manually.

## Running the application

```bash
cd cmd/defi-tui
go run main.go
```

On first launch, select "Configuration" from the main menu to set up your environment. You can update these settings at any time.

## Navigation

- Use arrow keys to navigate
- Tab to move between fields
- Enter to select/submit
- ESC to go back/exit

## Features

1. **Mint Tokens**
   - Enter amount
   - Click "Mint" or press Enter

2. **Lending Pool**
   - Enter amount
   - Choose "Deposit" or "Withdraw"

3. **NFT Management**
   - Enter collateral amount
   - Click "Mint NFT"

4. **Balance View**
   - Shows current token balance
   - Shows lending pool deposits
   - Shows owned NFTs

## Development


### Parallel Handling & Goroutines

This project uses Go's goroutines and WaitGroups for high-throughput parallel processing in several areas:

- **Contract Initialization:** All smart contract clients are initialized in parallel for faster startup.
- **Balance Updates:** Token, pool, and NFT balances are fetched concurrently for responsive UI updates.
- **Batch Transactions:** You can process multiple transactions (mint, deposit, withdraw, NFT mint) in parallel using the provided batch API.
- **Bulk NFT Minting:** Multiple NFTs can be minted simultaneously for speed and scalability.

These improvements allow the TUI to handle more requests efficiently and provide a smoother user experience. Make sure contract bindings are generated with `abigen` for these features to work.

The project structure:

```
├── cmd/
│   └── defi-tui/
│       └── main.go
├── internal/
│   ├── eth/
│   │   └── client.go
│   ├── config/
│   │   └── vault.go
│   └── ui/
│       ├── tui.go
│       ├── handlers.go
│       └── config.go
└── go.mod
```
