package mock

import (
	"math/big"

	"defi-tokenization-prototype/internal/eth/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Mock implementations for testing
type mockStableToken struct {
	client  *ethclient.Client
	address common.Address
}

type mockLendingPool struct {
	client  *ethclient.Client
	address common.Address
}

type mockNFTContract struct {
	client  *ethclient.Client
	address common.Address
}

func CreateMockStableToken(address common.Address, client *ethclient.Client) (contracts.StableToken, error) {
	return &mockStableToken{client: client, address: address}, nil
}

func CreateMockLendingPool(address common.Address, client *ethclient.Client) (contracts.LendingPool, error) {
	return &mockLendingPool{client: client, address: address}, nil
}

func CreateMockCollateralNFT(address common.Address, client *ethclient.Client) (contracts.CollateralNFT, error) {
	return &mockNFTContract{client: client, address: address}, nil
}

// Mock implementations of interface methods
func (m *mockStableToken) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

func (m *mockStableToken) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (m *mockLendingPool) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

func (m *mockLendingPool) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

func (m *mockLendingPool) GetUserBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (m *mockNFTContract) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

func (m *mockNFTContract) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	return []*big.Int{}, nil
}
