package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/99designs/keyring"
)

type Config struct {
	InfuraKey    string `json:"infura_key"`
	KeystorePath string `json:"keystore_path"`
	KeystorePass string `json:"keystore_pass"`
	TokenAddress string `json:"token_address"`
	PoolAddress  string `json:"pool_address"`
	NFTAddress   string `json:"nft_address"`
}

type Vault struct {
	ring    keyring.Keyring
	appName string
}

func NewVault() (*Vault, error) {
	appName := "defi-tui"

	// Try system keyring first
	ring, err := keyring.Open(keyring.Config{
		ServiceName: appName,
		AllowedBackends: []keyring.BackendType{
			keyring.SecretServiceBackend, // Linux
			keyring.KeychainBackend,      // macOS
			keyring.WinCredBackend,       // Windows
		},
	})

	if err != nil {
		// Fallback to file-based storage
		configDir, err := os.UserConfigDir()
		if err != nil {
			return nil, err
		}

		ring, err = keyring.Open(keyring.Config{
			ServiceName: appName,
			FileDir:     filepath.Join(configDir, "defi-tui"),
			FilePasswordFunc: func(string) (string, error) {
				return "default-password", nil // You might want to prompt for this
			},
			AllowedBackends: []keyring.BackendType{keyring.FileBackend},
		})
		if err != nil {
			return nil, err
		}
	}

	return &Vault{
		ring:    ring,
		appName: appName,
	}, nil
}

func (v *Vault) SaveConfig(cfg *Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return v.ring.Set(keyring.Item{
		Key:  "config",
		Data: data,
	})
}

func (v *Vault) LoadConfig() (*Config, error) {
	item, err := v.ring.Get("config")
	if err != nil {
		if err == keyring.ErrKeyNotFound {
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(item.Data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (v *Vault) SetEnvFromConfig() error {
	cfg, err := v.LoadConfig()
	if err != nil {
		return err
	}

	os.Setenv("INFURA_KEY", cfg.InfuraKey)
	os.Setenv("KEYSTORE_PATH", cfg.KeystorePath)
	os.Setenv("KEYSTORE_PASS", cfg.KeystorePass)
	os.Setenv("TOKEN_ADDRESS", cfg.TokenAddress)
	os.Setenv("POOL_ADDRESS", cfg.PoolAddress)
	os.Setenv("NFT_ADDRESS", cfg.NFTAddress)

	return nil
}
