package blockchain

import (
	"fmt"
	"log"
	"smart-rental/global"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitEthClient() {
    // Replace with your Infura project endpoint
    infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", global.Config.Infura.APIKey)

    // Connect to Infura
    client, err := ethclient.Dial(infuraURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
       
    }

    global.EtherClient = client
}
