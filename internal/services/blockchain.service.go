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
	SignMContractOnChain(privateKeyHex string, req requests.SignMContractOnChainReq) (string, error)
	PayDepositOnChain(privateKeyHex string, ContractId int64) (string, error)
	InputMeterReadingOnChain(privateKeyHex string, ContractId int64) (string, error)
	CreateBillOnChain(privateKeyHex string, ContractId int64) (string, error)
	PayBillOnChain(privateKeyHex string, ContractId int64) (string, error)
	CreateReturnRequestOnChain(privateKeyHex string, ContractId int64) (string, error)
	ApproveReturnRequestOnChain(privateKeyHex string, ContractId int64) (string, error)
}