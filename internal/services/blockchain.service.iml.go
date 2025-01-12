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
		contractManagementAddress: common.HexToAddress(global.Config.SmartContract.ContractManagement),
	}
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

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

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

func (b *BlockchainServiceImpl) CreateMContractOnChain(privateKeyHex string, req requests.CreateMContractOnChainReq) (string, error) {
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

	adjustedGasPrice := new(big.Int).Mul(gasPrice, big.NewInt(2)) // adjustedGasPrice = gasPrice * 2

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = 3000000
	auth.GasPrice = adjustedGasPrice

	// owner := crypto.PubkeyToAddress(privateKey.PublicKey)
	roomContract, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create room contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := roomContract.CreateMContract(
		auth,
		big.NewInt(req.ContractId),         // Contract ID
		req.ContractCode,                   // Contract code
		big.NewInt(req.LandlordId),         // Landlord ID
		big.NewInt(req.TenantId),           // Tenant ID
		big.NewInt(req.RoomId),             // Room ID
		big.NewInt(req.ActualPrice),        // Actual price
		big.NewInt(req.Deposit),            // Deposit
		big.NewInt(req.BeginDate),          // Begin date
		big.NewInt(req.EndDate),            // End date
		req.PaymentMethod,                  // Payment method
		req.ElectricityMethod,              // Electricity method
		big.NewInt(req.ElectricityCost),    // Electricity cost
		req.WaterMethod,                    // Water method
		big.NewInt(req.WaterCost),          // Water cost
		big.NewInt(req.InternetCost),       // Internet cost
		big.NewInt(req.ParkingFee),         // Parking fee
		req.ResponsibilityA,                // Responsibility A
		req.ResponsibilityB,                // Responsibility B
		req.GeneralResponsibility,          // General responsibility
		req.SignatureA,                     // Signature A
		big.NewInt(req.SignedTimeA),        // Signed time A
		req.SignatureB,                     // Signature B
		big.NewInt(req.SignedTimeB),        // Signed time B
		big.NewInt(req.ContractTemplateId), // Contract template ID
	)

	if err != nil {
		return "", fmt.Errorf("failed to create room on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

// GetRoomFromBlockchain implements BlockchainService.
func (b *BlockchainServiceImpl) GetMContractByIDOnChain(contractId int64) (*responses.MContractOnChainRes, error) {
	roomContract, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create room contract instance: %w", err)
	}

	// Fetch room details
	bigContractId := big.NewInt(contractId)
	out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, out14, out15, out16, out17, out18, out19, out20, out21, out22, out23, out24, out25, out26, out27, out28, err := roomContract.GetMContract(&bind.CallOpts{}, bigContractId)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve room data from blockchain: %w", err)
	}

	// Format data for the response
	blockchainContract := &responses.MContractOnChainRes{
		ID:                    out0.Int64(),
		Code:                  out1,
		Landlord:              out2.Int64(),
		Tenant:                out3.Int64(),
		RoomID:                out4.Int64(),
		ActualPrice:           out5.Int64(),
		Deposit:               out6.Int64(),
		BeginDate:             out7.Int64(),
		EndDate:               out8.Int64(),
		PaymentMethod:         out9,
		ElectricityMethod:     out10,
		ElectricityCost:       out11.Int64(),
		WaterMethod:           out12,
		WaterCost:             out13.Int64(),
		InternetCost:          out14.Int64(),
		ParkingFee:            out15.Int64(),
		ResponsibilityA:       out16,
		ResponsibilityB:       out17,
		GeneralResponsibility: out18,
		SignatureA:            out19,
		SignedTimeA:           out20.Int64(),
		SignatureB:            out21,
		SignedTimeB:           out22.Int64(),
		ContractTemplateID:    out23.Int64(),
		PreRentalStatus:       uint8(out24),
		RentalProcessStatus:   uint8(out25),
		PostRentalStatus:      uint8(out26),
		CreatedAt:             out27.Int64(),
		UpdatedAt:             out28.Int64(),
	}

	return blockchainContract, nil
}

func (b *BlockchainServiceImpl) GetListMContractByStatus(contractIds []int32, statusID int64, userId int64, isLandlord bool) ([]gen.ContractManagementMContract, error) {
	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	bigUserId := big.NewInt(userId)
	ids := make([]*big.Int, len(contractIds))
	for i, id := range contractIds {
		ids[i] = big.NewInt(int64(id))
	}

	contracts, err := contractManagement.GetContractsByPreRentalStatus(&bind.CallOpts{}, uint8(statusID), ids, isLandlord, bigUserId)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve contracts data from blockchain: %w", err)
	}

	return contracts, nil
}

func (b *BlockchainServiceImpl) DeclineMContractOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	tx, err := contractManagement.DeclineContract(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to decline contract on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) SignMContractOnChain(privateKeyHex string, req requests.SignMContractOnChainReq) (string, error) {
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
	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.SignContractByTenant(
		auth,
		big.NewInt(req.ContractId),
		req.SignatureB,
	)

	if err != nil {
		return "", fmt.Errorf("failed to sign contract on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) PayDepositOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.PayDeposit(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to pay deposit on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) InputMeterReadingOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.InputMeterReading(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to input meter reading on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) CreateBillOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.CreateBill(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to create bill on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) PayBillOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.PayBill(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to pay bill on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) CreateReturnRequestOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.CreateReturnRequest(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to create return request on blockchain: %w", err)
	}

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (b *BlockchainServiceImpl) ApproveReturnRequestOnChain(privateKeyHex string, ContractId int64) (string, error) {
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

	contractManagement, err := gen.NewContractManagement(b.contractManagementAddress, b.client)
	if err != nil {
		return "", fmt.Errorf("failed to create contract instance: %w", err)
	}

	// Prepare parameters for the room creation
	tx, err := contractManagement.ApproveReturnRequest(
		auth,
		big.NewInt(ContractId),
	)

	if err != nil {
		return "", fmt.Errorf("failed to approve return request on blockchain: %w", err)
	}

	fmt.Print("approve return request on blockchain: %w", tx.Hash().Hex())

	fmt.Print("Transaction hash: " + tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}
