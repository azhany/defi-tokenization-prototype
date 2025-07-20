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

Set the following environment variables:

- `INFURA_KEY`: Your Infura project key
- `KEYSTORE_PATH`: Path to your Ethereum keystore file
- `KEYSTORE_PASS`: Password for your keystore
- `TOKEN_ADDRESS`: Deployed stable token contract address
- `POOL_ADDRESS`: Deployed lending pool contract address
- `NFT_ADDRESS`: Deployed NFT contract address

## Running the application

```bash
cd cmd/defi-tui
go run main.go
```

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

The project structure:

```
├── cmd/
│   └── defi-tui/
│       └── main.go
├── internal/
│   ├── eth/
│   │   └── client.go
│   └── ui/
│       ├── tui.go
│       └── handlers.go
└── go.mod
```
