package services

import (
	"context"
	"fmt"
	"math/big"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	room "smart-rental/pkg/blockchain/gen"
	c "smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type RoomServiceImpl struct {
	repo           *dataaccess.Queries
	storageService StorageSerivce
}

func NewRoomServiceImpl(storage StorageSerivce) RoomService {
	return &RoomServiceImpl{
		repo:           dataaccess.New(global.Db),
		storageService: storage,
	}
}

// CreateRoom implements RoomService.
func (r *RoomServiceImpl) CreateRoom(req requests.CreateRoomForm) *responses.ResponseData {

	// create new room
	if exist, _ := r.storageService.IsBucketExists(constants.BUCKET_NAME); !exist {
		err := r.storageService.CreateBucket(constants.BUCKET_NAME)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}
	var params dataaccess.CreateRoomParams
	c.MapStruct(req, &params)
	id, err := r.repo.CreateRoom(context.Background(), params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	// Add to blockchain
	chainID, _ := global.EtherClient.NetworkID(context.Background())
	privateKeyHex := "e5a0d26cd7866afbea195c376b75a76d86d65e458c25f4702c46f1378ea0ce42"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	gasPrice, err := global.EtherClient.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}
	contractAddress := common.HexToAddress("0x36a38c0aa01433d3732260b9706a1fB07aFa30e5")
	roomContract, err := room.NewListingContract(contractAddress, global.EtherClient)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	auth.GasLimit = 3000000
	auth.GasPrice = gasPrice

	//de
	roomID := big.NewInt(int64(id))
	owner := crypto.PubkeyToAddress(privateKey.PublicKey)
	totalPrice := big.NewInt(int64(*req.TotalPrice))
	deposit := big.NewInt(int64(req.Deposit))
	status := big.NewInt(int64(req.Status))
	isRent := req.IsRent

	tx, err := roomContract.CreateRoom(auth, roomID, owner.Big(), totalPrice, deposit, status, isRent)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	fmt.Println("Room created on blockchain. Transaction hash:", tx.Hash().Hex())
	// update images url
	var urls []string
	for _, fileName := range req.RoomImages {
		f, _ := fileName.Open()
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		roomID := fmt.Sprintf("room_%d", id)
		objKey := fmt.Sprintf("%s/%s/%d%s", constants.ROOM_OBJ, roomID, timestamp, fileExt)

		url, err := r.storageService.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		urls = append(urls, url)

	}

	var updateRoom dataaccess.UpdateRoomParams
	updateRoom.ID = id
	updateRoom.RoomImages = urls
	_, updateErr := r.repo.UpdateRoom(context.Background(), updateRoom)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// GetRooms implements RoomService.
func (r *RoomServiceImpl) GetRooms() *responses.ResponseData {
	rooms, err := r.repo.GetRooms(context.Background())

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       rooms,
	}
}

// GetRoomByID implements RoomService.
func (r *RoomServiceImpl) GetRoomByID(id int) *responses.ResponseData {

	roomData, err := r.repo.GetRoomByID(context.Background(), int32(id))
	// // Connect to Ethereum blockchain and fetch room data from smart contract
	// contractAddress := common.HexToAddress("0x36a38c0aa01433d3732260b9706a1fB07aFa30e5") // Replace with your deployed contract address
	// roomContract, err := room.NewRoom(contractAddress, global.EtherClient)
	if err != nil {
		if (roomData.ID == 0) {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    "Phòng không tồn tại",
				Data:       false,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	// // Fetch room details from the blockchain
	// // Fetch room details from the blockchain
	// bigID := big.NewInt(int64(id))
	// out0, out1, out2, out3, out4, out5, out6, out7, err := roomContract.GetRoom(&bind.CallOpts{}, bigID)
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Failed to retrieve room data from blockchain: " + err.Error(),
	// 		Data:       nil,
	// 	}
	// }

	// // Combine data from the database and blockchain (you may customize this as per your requirements)
	// roomData := map[string]interface{}{
	// 	"blockchain_room": map[string]interface{}{
	// 		"id":         out0,
	// 		"owner":      out1,
	// 		"totalPrice": out2,
	// 		"deposit":    out3,
	// 		"status":     out4.String(),
	// 		"isRent":     out5,
	// 		"createdAt":  out6.String(),
	// 		"updatedAt":  out7.String(),
	// 	},
	// }

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       roomData,
	}
}

// GetRoomByID implements RoomService.
// func (r *RoomServiceImpl) GetRoomByID(id int) *responses.ResponseData {
// 	room, err := r.repo.GetRoomByID(context.Background(), int32(id))

// 	if err != nil {
// 		return &responses.ResponseData{
// 			StatusCode: http.StatusInternalServerError,
// 			Message:    err.Error(),
// 			Data:       nil,
// 		}
// 	}
// 	return &responses.ResponseData{
// 		StatusCode: http.StatusOK,
// 		Message:    responses.StatusSuccess,
// 		Data:       room,
// 	}
// }

// SearchRoomByAddress implements RoomService.
func (r *RoomServiceImpl) SearchRoomByAddress(address string) *responses.ResponseData {
	rooms, err := r.repo.SearchRoomByAddress(context.Background(), &address)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       rooms,
	}
}

// LikeRoom implements RoomService.
func (r *RoomServiceImpl) LikeRoom(roomID int, userID int) *responses.ResponseData {
	param := dataaccess.CheckUserLikedRoomParams{
		RoomID: int32(roomID),
		UserID: int32(userID),
	}
	_, err := r.repo.CheckUserLikedRoom(context.Background(), param)
	var isLiked bool

	if err != nil {
		isLiked = true
		param := dataaccess.LikeRoomParams{
			RoomID: int32(roomID),
			UserID: int32(userID),
		}
		err := r.repo.LikeRoom(context.Background(), param)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	} else {
		isLiked = false
		param := dataaccess.UnlikeRoomParams{
			RoomID: int32(roomID),
			UserID: int32(userID),
		}
		err := r.repo.UnlikeRoom(context.Background(), param)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}
	res := map[string]bool{"isLiked": isLiked}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       res,
	}
}

// GetLikedRooms implements RoomService.
func (r *RoomServiceImpl) GetLikedRooms(userID int) *responses.ResponseData {
	rooms, err := r.repo.GetLikedRooms(context.Background(), int32(userID))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       rooms,
	}
}

// GetRoomByStatus implements RoomService.
func (r *RoomServiceImpl) GetRoomByStatus(status int) *responses.ResponseData {
	rooms, err := r.repo.GetRoomsByStatus(context.Background(), int32(status))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       rooms,
	}
}

// UpdateRoom implements RoomService.
func (r *RoomServiceImpl) UpdateRoom(req requests.UpdateRoomRequest) *responses.ResponseData {
	// delete file from s3
	for _, urlString := range req.DeleteFiles {
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		objKey := parsedURL.Path

		objKey = strings.TrimPrefix(objKey, "/")
		deleteErr := r.storageService.DeleteObject(constants.BUCKET_NAME, objKey)

		if deleteErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to delete object from s3",
				Data:       nil,
			}
		}
	}

	// update file to s3
	var urls []string
	for _, fileName := range req.RoomImages {
		f, _ := fileName.Open()
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		roomID := fmt.Sprintf("room_%d", req.ID)
		objKey := fmt.Sprintf("%s/%s/%d%s", constants.ROOM_OBJ, roomID, timestamp, fileExt)

		url, err := r.storageService.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		urls = append(urls, url)

	}

	// update data in db
	var param dataaccess.UpdateRoomParams
	c.MapStruct(req, &param)
	param.RoomImages = urls

	_, updateErr := r.repo.UpdateRoom(context.Background(), param)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}
