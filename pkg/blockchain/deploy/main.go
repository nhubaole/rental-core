package main

import (
	"context"
	"fmt"
	"math/big"
	contract "smart-rental/pkg/blockchain/gen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	api := "5f4a92e00828480ea559e5f4580861cd"
	var infuraURL = fmt.Sprintf("https://sepolia.infura.io/v3/%s", api)
	
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		panic(err)
	}

	// b, err := os.ReadFile("pkg/blockchain/keystore/UTC--2024-11-04T08-44-39.616334100Z--9f00e9040ce36ffd0c4829e45a2487030966f2a7")
	// if err != nil {
	// 	panic(err)
	// }
	// key, err := keystore.DecryptKey(b, "F42C765A563F2D7BA8ED57AECD8B6")
	privateKeyHex := "2da01a1b767463db18eabaf32a93bf8b0b8816e2b01a68e0c8e5e3baac2f2e61"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err)
	}
	
	if err != nil {
		panic(err)
	}
	fmt.Printf("Private key: 0x%x\n", privateKey.D)

	add := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("====", add)
	nonce, err := client.PendingNonceAt(context.Background(), add)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	a, tx, _, err := contract.DeployContractManagement(auth, client) //contract.DeployListingContract(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println("-----------------------------------")
	fmt.Println(a.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println("-----------------------------------")
}