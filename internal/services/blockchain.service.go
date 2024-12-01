package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type BlockchainService interface {
	GetRoomByIDOnChain(roomID int64)(*responses.RoomOnChainRes, error)
	CreateRoomOnBlockchain(privateKeyHex string, req requests.CreateRoomOnChainReq) (string, error)

	GetMContractByIDOnChain(roomID int64)(*responses.MContractOnChainRes, error)
	CreateMContractOnChain(privateKeyHex string, req requests.CreateMContractOnChainReq) (string, error)

	GetContractByIDOnChain(contractID int64)(*responses.ContractOnChainRes, error)
	CreateContractOnBlockchain(privateKeyHex string, req requests.CreateContractOnChainReq) (string, error)
	// CreateLeaseAgreementProducerContract(privateKeyHex string, req requests.CreateLeaseAgreementOnChainReq,) (string, error)
	// GetAllContractsOnChain(participantAddress string)([]responses.ContractOnChainRes, error)
}