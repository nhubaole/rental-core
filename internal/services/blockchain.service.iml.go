package services

import (
	"context"
	"fmt"
	"math/big"
	"smart-rental/global"
	gen "smart-rental/pkg/blockchain/gen"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainServiceImpl struct {
	client                    *ethclient.Client
	listingContractAddress    common.Address
	lapcAddress               common.Address
	contractManagementAddress common.Address
}

func NewBlockchainServiceImpl() BlockchainService {
	return &BlockchainServiceImpl{
		client:                    global.EtherClient,
		listingContractAddress:    common.HexToAddress(global.Config.SmartContract.ListingContractAddress),
		lapcAddress:               common.HexToAddress(global.Config.SmartContract.LeaseAgreementProducerContract),
		contractManagementAddress: common.HexToAddress(global.Config.SmartContract.LeaseAgreementProducerContract),
	}
}

// // GetAllContractsOnChain implements BlockchainService.
// func (b *BlockchainServiceImpl) GetAllContractsOnChain(participantAddress string) ([]responses.ContractOnChainRes, error) {
// 	// Tạo một instance của hợp đồng LeaseAgreementProducerContract
// 	leaseContract, err := room.NewLeaseAgreementProducerContract(b.lapcAddress, b.client)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create lease agreement producer contract instance: %w", err)
// 	}

// 	// Gọi hàm getContractsByParticipant từ smart contract
// 	participant := common.HexToAddress(participantAddress)
// 	contractAddresses, err := leaseContract.GetContractsByParticipant(&bind.CallOpts{}, participant)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve contracts from blockchain: %w", err)
// 	}

// 	// Tạo danh sách kết quả để trả về
// 	var contracts []responses.ContractOnChainRes
// 	for _, contractAddress := range contractAddresses {
// 		// Lấy chi tiết từng hợp đồng bằng hàm getLeaseContractDetails
// 		contractDetails, err := leaseContract.GetLeaseContractDetails(&bind.CallOpts{}, contractAddress)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to retrieve contract details for address %s: %w", contractAddress.Hex(), err)
// 		}

// 		// Thêm thông tin hợp đồng vào danh sách
// 		contracts = append(contracts, responses.ContractOnChainRes{
// 			ContractAddress: contractAddress.Hex(),
// 			Landlord:        contractDetails.Landlord.Hex(),
// 			Tenant:          contractDetails.Tenant.Hex(),
// 			RoomID:          contractDetails.RoomId.Int64(),
// 			ActualPrice:     contractDetails.ActualPrice.Int64(),
// 			DepositAmount:   contractDetails.DepositAmount.Int64(),
// 			BeginDate:       contractDetails.BeginDate.Int64(),
// 			EndDate:         contractDetails.EndDate.Int64(),
// 			ContractCode:    contractDetails.ContractCode,
// 		})
// 	}

// 	return contracts, nil
// }

// // CreateLeaseAgreementProducerContract creates a new lease agreement on the blockchain
// func (b *BlockchainServiceImpl) CreateLeaseAgreementProducerContract(
// 	privateKeyHex string,
// 	req requests.CreateLeaseAgreementOnChainReq,
// ) (string, error) {
// 	// Get the chain ID for the network
// 	chainID, err := b.client.NetworkID(context.Background())
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get network ID: %w", err)
// 	}

// 	// Parse the private key
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to parse private key: %w", err)
// 	}

// 	// Suggest gas price
// 	gasPrice, err := b.client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return "", fmt.Errorf("failed to suggest gas price: %w", err)
// 	}

// 	// Set up transaction options with the private key, gas price, and limit
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create transactor: %w", err)
// 	}
// 	auth.GasLimit = 3000000
// 	auth.GasPrice = gasPrice

// 	// Create a new instance of the LeaseAgreementProducerContract
// 	leaseContract, err := room.NewLeaseAgreementProducerContract(b.lapcAddress, b.client)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create lease agreement producer contract instance: %w", err)
// 	}

// 	// Call createLeaseContract with the necessary parameters
// 	tx, err := leaseContract.CreateLeaseContract(
// 		auth,
// 		common.HexToAddress(req.TenantAddress), // Tenant's address
// 		big.NewInt(req.RoomID),                 // Room ID
// 		big.NewInt(int64(req.ActualPrice)),     // Actual price
// 		big.NewInt(int64(req.DepositAmount)),   // Deposit amount
// 		big.NewInt(req.BeginDate),              // Begin date
// 		big.NewInt(req.EndDate),                // End date
// 		req.ContractCode,                       // Contract code
// 		req.SignatureA,                         // Landlord's signature
// 		big.NewInt(req.SignedTimeA),            // Landlord's signing timestamp
// 		req.PaymentMethod,                      // Payment method
// 		req.ElectricityMethod,                  // Electricity method
// 		big.NewInt(req.ElectricityCost),        // Electricity cost
// 		req.WaterMethod,                        // Water method
// 		big.NewInt(req.WaterCost),              // Water cost
// 		big.NewInt(req.InternetCost),           // Internet cost
// 		big.NewInt(req.ParkingFee),             // Parking fee
// 		req.ResponsibilityA,                    // Landlord's responsibilities
// 		req.ResponsibilityB,                    // Tenant's responsibilities
// 		req.GeneralResponsibility,              // General responsibilities
// 		big.NewInt(req.ContractTemplateID),     // Contract template ID
// 	)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create lease agreement on blockchain: %w", err)
// 	}

// 	fmt.Println(tx.Hash())
// 	// Return the transaction hash
// 	return tx.Hash().Hex(), nil
// }

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
	roomContract, err := gen.NewListingContract(b.listingContractAddress, b.client)
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
	roomContract, err := gen.NewListingContract(b.listingContractAddress, b.client)
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

type CreateLeaseContractParams struct {
	Code                  string
	Landlord              common.Address
	Tenant                common.Address
	RoomID                *big.Int
	ActualPrice           *big.Int
	DepositAmount         *big.Int
	BeginDate             *big.Int
	EndDate               *big.Int
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       *big.Int
	WaterMethod           string
	WaterCost             *big.Int
	InternetCost          *big.Int
	ParkingFee            *big.Int
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           *big.Int
	ContractTemplateId    *big.Int
}

func callCreateLeaseContract(auth *bind.TransactOpts, contract *gen.LeaseContractManagement, params CreateLeaseContractParams) (*types.Transaction, error) {
	tx, err := contract.CreateLeaseContract(
		auth,
		params.Code,
		params.Landlord,
		params.Tenant,
		params.RoomID,
		params.ActualPrice,
		params.DepositAmount,
		params.BeginDate,
		params.EndDate,
		params.PaymentMethod,
		params.ElectricityMethod,
		params.ElectricityCost,
		params.WaterMethod,
		params.WaterCost,
		params.InternetCost,
		params.ParkingFee,
		params.ResponsibilityA,
		params.ResponsibilityB,
		params.GeneralResponsibility,
		params.SignatureA,
		params.SignedTimeA,
		params.ContractTemplateId,
	)
	return tx, err
}

// CreateContractOnBlockchain implements BlockchainService.
func (b *BlockchainServiceImpl) CreateContractOnBlockchain(privateKeyHex string, req requests.CreateContractOnChainReq) (string, error) {
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

	// owner := crypto.PubkeyToAddress(privateKey.PublicKey)
	contractManagement, err := gen.NewLeaseContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	params := CreateLeaseContractParams{
		Code:                  "LEASE-123",
		Landlord:              common.HexToAddress("9F00E9040CE36fFD0c4829E45a2487030966f2A7"),
		Tenant:                common.HexToAddress("ae4228c40e0409E0fD61af9d41906ABE30072964"),
		RoomID:                big.NewInt(101),
		ActualPrice:           big.NewInt(100000),
		DepositAmount:         big.NewInt(20000),
		BeginDate:             big.NewInt(1700000000), // UNIX timestamp
		EndDate:               big.NewInt(1730000000), // UNIX timestamp
		PaymentMethod:         "Bank Transfer",
		ElectricityMethod:     "Prepaid",
		ElectricityCost:       big.NewInt(5000),
		WaterMethod:           "Monthly",
		WaterCost:             big.NewInt(2000),
		InternetCost:          big.NewInt(1500),
		ParkingFee:            big.NewInt(500),
		ResponsibilityA:       "Landlord's responsibilities",
		ResponsibilityB:       "Tenant's responsibilities",
		GeneralResponsibility: "General rules",
		SignatureA:            "LandlordSignature",
		SignedTimeA:           big.NewInt(1700000100), // UNIX timestamp
		ContractTemplateId:    big.NewInt(1),
	}

	tx, err := callCreateLeaseContract(auth, contractManagement, params)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction sent, hash:", tx.Hash().Hex())
	}

	return tx.Hash().Hex(), nil
}

// GetContractByIDOnChain implements BlockchainService.
func (b *BlockchainServiceImpl) GetContractByIDOnChain(contractID int64) (*responses.ContractOnChainRes, error) {
	contractManagement, err := gen.NewLeaseContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create room contract instance: %w", err)
	}

	// Fetch room details
	bigContractID := big.NewInt(contractID)
	out0, err := contractManagement.GetContractById(&bind.CallOpts{}, bigContractID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve room data from blockchain: %w", err)
	}

	// Format data for the response
	blockchainRoom := &responses.ContractOnChainRes{
		ID:                int32(out0.ContractId.Int64()),
		Code:              out0.ContractCode,
		PartyA:            123,
		PartyB:            456,
		RequestID:         789,
		RoomID:            101,
		ActualPrice:       1500.75,
		PaymentMethod:     out0.PaymentMethod,
		ElectricityMethod: "Monthly",
		ElectricityCost:   300.00,
		WaterMethod:       "Quarterly",
		WaterCost:         150.00,
		InternetCost:      100.00,
		ParkingFee:        float64(out0.ParkingFee.Int64()),
		Deposit:           500.00,
		BeginDate:         out0.BeginDate.Int64(),
		EndDate:           out0.EndDate.Int64(),
		ResponsibilityA:   "Maintain property",
		ResponsibilityB:   "Pay rent on time",
		GeneralResponsibility: out0.GeneralResponsibility,
		SignatureA:        "Signature123",
		SignedTimeA:       out0.SignedTimeA.Int64(),
		SignatureB:        out0.SignatureB,
		SignedTimeB:       out0.SignedTimeB.Int64(),
		CreatedAt:         out0.CreatedAt.Int64(),
		UpdatedAt:         out0.UpdatedAt.Int64(),
		ContractTemplateID: int32(out0.ContractTemplateId.Int64()),
	}
	return blockchainRoom, nil
}