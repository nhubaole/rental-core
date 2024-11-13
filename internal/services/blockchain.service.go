package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type BlockchainService interface {
	GetRoomFromBlockchain(roomID int64)(*responses.RoomResOnChain, error)
	CreateRoomOnBlockchain(privateKeyHex string, req requests.CreateRoomOnChainReq) (string, error)
	CreateLeaseAgreementProducerContract() error
}