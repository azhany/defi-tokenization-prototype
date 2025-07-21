package contracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// StableToken interface defines the methods for interacting with the stable token contract
type StableToken interface {
	Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

// LendingPool interface defines the methods for interacting with the lending pool contract
type LendingPool interface {
	Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
	Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
	GetUserBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error)
}

// CollateralNFT interface defines the methods for interacting with the NFT contract
type CollateralNFT interface {
	Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
	TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error)
}
