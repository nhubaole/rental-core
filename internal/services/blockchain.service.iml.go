package services

import (
	"context"
	"fmt"
	"math/big"
	"smart-rental/global"
	room "smart-rental/pkg/blockchain/gen"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainServiceImpl struct {
	client                 *ethclient.Client
	listingContractAddress common.Address
	lapcAddress common.Address
}



func NewBlockchainServiceImpl() BlockchainService {
	return &BlockchainServiceImpl{
		client:                 global.EtherClient,
		listingContractAddress: common.HexToAddress(global.Config.SmartContract.ListingContractAddress),
		lapcAddress: common.HexToAddress(global.Config.SmartContract.LeaseAgreementProducerContract),
	}
}

// CreateLeaseAgreementProducerContract creates a new lease agreement on the blockchain
func (b *BlockchainServiceImpl) CreateLeaseAgreementProducerContract(
	privateKeyHex string,
	req requests.CreateLeaseAgreementOnChainReq,
) (string, error) {
	// Get the chain ID for the network
	chainID, err := b.client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get network ID: %w", err)
	}

	// Parse the private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	// Suggest gas price
	gasPrice, err := b.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %w", err)
	}

	// Set up transaction options with the private key, gas price, and limit
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice = gasPrice

	// Create a new instance of the LeaseAgreementProducerContract
	leaseContract, err := room.NewLeaseAgreementProducerContract(b.lapcAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create lease agreement producer contract instance: %w", err)
	}

	// Call createLeaseContract with the necessary parameters
	tx, err := leaseContract.CreateLeaseContract(
		auth,
		common.HexToAddress(req.TenantAddress),    // Tenant's address
		big.NewInt(req.RoomID),                    // Room ID
		big.NewInt(int64(req.ActualPrice)),        // Actual price
		big.NewInt(int64(req.DepositAmount)),      // Deposit amount
		big.NewInt(req.BeginDate),                 // Begin date
		big.NewInt(req.EndDate),                   // End date
		req.ContractCode,                          // Contract code
		req.SignatureA,                            // Landlord's signature
		big.NewInt(req.SignedTimeA),               // Landlord's signing timestamp
		req.PaymentMethod,                         // Payment method
		req.ElectricityMethod,                     // Electricity method
		big.NewInt(req.ElectricityCost),           // Electricity cost
		req.WaterMethod,                           // Water method
		big.NewInt(req.WaterCost),                 // Water cost
		big.NewInt(req.InternetCost),              // Internet cost
		big.NewInt(req.ParkingFee),                // Parking fee
		req.ResponsibilityA,                       // Landlord's responsibilities
		req.ResponsibilityB,                       // Tenant's responsibilities
		req.GeneralResponsibility,                 // General responsibilities
		big.NewInt(req.ContractTemplateID),        // Contract template ID
	)
	if err != nil {
		return "", fmt.Errorf("failed to create lease agreement on blockchain: %w", err)
	}

	// Return the transaction hash
	return tx.Hash().Hex(), nil
}


// CreateRoomOnBlockchain implements BlockchainService.
func (b *BlockchainServiceImpl) CreateRoomOnBlockchain(privateKeyHex string, req requests.CreateRoomOnChainReq) (string, error) {
	chainID, err := b.client.NetworkID(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get network ID: %w", err)
	}

	// Parse the private key provided by the user
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	gasPrice, err := b.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice = gasPrice

	owner := crypto.PubkeyToAddress(privateKey.PublicKey)
	roomContract, err := room.NewListingContract(b.listingContractAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create room contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := roomContract.CreateRoom(
		auth,
		big.NewInt(req.RoomID),
		owner.Big(),
		big.NewInt(int64(req.TotalPrice)),
		big.NewInt(req.Deposit),
		big.NewInt(req.Deposit),
		req.IsRent,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create room on blockchain: %w", err)
	}

	return tx.Hash().Hex(), nil
}

// GetRoomFromBlockchain implements BlockchainService.
func (b *BlockchainServiceImpl) GetRoomByIDOnChain(roomID int64) (*responses.RoomOnChainRes, error) {
	roomContract, err := room.NewListingContract(b.listingContractAddress, b.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create room contract instance: %w", err)
	}

	// Fetch room details
	bigRoomID := big.NewInt(roomID)
	out0, _, out2, out3, out4, out5, out6, _, err := roomContract.GetRoom(&bind.CallOpts{}, bigRoomID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve room data from blockchain: %w", err)
	}

	// Format data for the response
	blockchainRoom := &responses.RoomOnChainRes{
		ID:         out0.Int64(),
		TotalPrice: int(out2.Int64()),
		Deposit:    out3.Int64(),
		Status:     out4.Int64(),
		IsRent:     out5,
		CreatedAt:  out6.Int64(),
	}
	return blockchainRoom, nil
}
