package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"
	"smart-rental/global"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func CreateWallet(phoneNumber string) (string, string, error) {
	// Define the keystore folder to store wallets
	keystoreDir := "./keystore" // Ensure this folder is secured in production
	ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	// Password to encrypt the private key (for demo purposes, use stronger mechanisms in production)
	password := global.Config.JWT.SecretKey

	// Create a new account
	account, err := ks.NewAccount(password)
	if err != nil {
		return "", "", err
	}

	// Save the wallet details (you might want to store these in a more secure database)
	return account.Address.Hex(), filepath.Join(keystoreDir, account.URL.Path), nil
}

func GetWalletBalance(walletAddress string) (*big.Int, error) {

	address := common.HexToAddress(walletAddress)

	// Get the balance
	balance, err := global.EtherClient.BalanceAt(context.Background(), address, nil) // nil for latest block
	if err != nil {
		fmt.Printf("Failed to get balance: %v", err)
		return nil, err
	}

	return balance, nil
}
