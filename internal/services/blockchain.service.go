package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type BlockchainService interface {
	GetRoomByIDOnChain(roomID int64)(*responses.RoomOnChainRes, error)
	CreateRoomOnBlockchain(privateKeyHex string, req requests.CreateRoomOnChainReq) (string, error)
	CreateLeaseAgreementProducerContract(privateKeyHex string, req requests.CreateLeaseAgreementOnChainReq,) (string, error)
	GetAllContractsOnChain(participantAddress string)([]responses.ContractOnChainRes, error)
}