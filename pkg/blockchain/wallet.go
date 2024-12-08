package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"crypto/ecdsa"
	"smart-rental/global"
	"encoding/hex"
	"errors"
	"crypto/rand"

	
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func CreateWallet(phoneNumber string) (string, string, error) {
	// Generate a new private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	// Convert the private key to a hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Derive the public key and Ethereum address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("cannot cast public key to ECDSA")
	}

	// Generate the Ethereum address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// Return the private key (hex) and Ethereum address
	return privateKeyHex, address, nil
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
