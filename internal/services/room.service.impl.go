package services

import (
	"context"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	c "smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strings"
	"time"
)

type RoomServiceImpl struct {
	repo              *dataaccess.Queries
	storageService    StorageSerivce
	blockchainService BlockchainService
}


func NewRoomServiceImpl(storage StorageSerivce, blockchain BlockchainService) RoomService {
	return &RoomServiceImpl{
		repo:              dataaccess.New(global.Db),
		storageService:    storage,
		blockchainService: blockchain,
	}
}

// GetRoomByOwner implements RoomService.
func (r *RoomServiceImpl) GetRoomByOwner(userID int) *responses.ResponseData {
	rooms, err := r.repo.GetRoomsByOwner(context.Background(), int32(userID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data: map[string]interface{}{
			"count": len(rooms),
			"rooms": rooms,
		},
	}

}

// CreateRoom implements RoomService.
func (r *RoomServiceImpl) CreateRoom(req requests.CreateRoomForm, userID int) *responses.ResponseData {
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

	user, err := r.repo.GetUserByID(context.Background(), int32(userID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	// Add to blockchain
	// privateKeyHex := "e5a0d26cd7866afbea195c376b75a76d86d65e458c25f4702c46f1378ea0ce42"
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       false,
	// 	}
	// }
	paramsOnChain := &requests.CreateRoomOnChainReq{
		RoomID:     int64(id),
		TotalPrice: int(*req.TotalPrice),
		Deposit:    int64(req.Deposit),
		Status:     int64(req.Status),
		IsRent:     req.IsRent,
	}

	if _, err := r.blockchainService.CreateRoomOnBlockchain(*user.PrivateKeyHex, *paramsOnChain); err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

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

	room, _ := r.repo.GetRoomByID(context.Background(), int32(id))
	var updateRoom dataaccess.UpdateRoomParams
	common.MapStruct(room, &updateRoom)
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
		Data:       id,
	}
}

// GetRooms implements RoomService.
func (r *RoomServiceImpl) GetRooms(userID int) *responses.ResponseData {
	rooms, err := r.repo.GetRooms(context.Background(), int32(userID))

	if len(rooms) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    responses.StatusNoData,
			Data:       []dataaccess.GetRoomsRow{},
		}
	}

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
		Data: map[string]interface{}{
			"count": len(rooms),
			"rooms": rooms,
		},
	}
}

// GetRoomByID implements RoomService.
func (r *RoomServiceImpl) GetRoomByID(id int) *responses.ResponseData {
	// Fetch from database
	roomData, err := r.repo.GetRoomByID(context.Background(), int32(id))
	if err != nil {
		if roomData.ID == 0 {
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

	//Fetch room details from the blockchain
	roomOnChain, err := r.blockchainService.GetRoomByIDOnChain(int64(id))
	if err != nil {
		if roomData.ID == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}

	// Override attribute on chain to response
	roomData.TotalPrice = common.IntToFloat64Ptr(roomOnChain.TotalPrice)
	roomData.Deposit = float64(roomOnChain.Deposit)
	roomData.Status = int32(roomOnChain.Status)
	roomData.IsRent = roomOnChain.IsRent
	roomData.CreatedAt = c.Int64ToPgTimestamptz(roomOnChain.CreatedAt, true)
	roomData.UpdatedAt = c.Int64ToPgTimestamptz(roomOnChain.UpdatedAt, true)

	owner, err := r.repo.GetUserByID(context.Background(), roomData.Owner)
	if err != nil {
		if roomData.ID == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}
	var listRoomNumberJson map[int]interface{}
	json.Unmarshal([]byte(roomData.ListRoomNumbers), &listRoomNumberJson)
	var result = responses.GetRoomByIDRes{
		ID:              roomData.ID,
		Title:           roomData.Title,
		Address:         roomData.Address,
		RoomNumber:      roomData.RoomNumber,
		RoomImages:      roomData.RoomImages,
		Utilities:       roomData.Utilities,
		Description:     roomData.Description,
		RoomType:        roomData.RoomType,
		AvailableFrom:   roomData.AvailableFrom,
		ListRoomNumbers: listRoomNumberJson,
		Owner:           owner,
		Capacity:        roomData.Capacity,
		Gender:          roomData.Gender,
		Area:            roomData.Area,
		TotalPrice:      roomData.TotalPrice,
		Deposit:         roomData.Deposit,
		ElectricityCost: roomData.ElectricityCost,
		WaterCost:       roomData.WaterCost,
		InternetCost:    roomData.InternetCost,
		IsParking:       roomData.IsParking,
		ParkingFee:      roomData.ParkingFee,
		Status:          roomData.Status,
		IsRent:          roomData.IsRent,
		CreatedAt:       roomData.CreatedAt,
		UpdatedAt:       roomData.UpdatedAt,
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// SearchRoomByAddress implements RoomService.
func (r *RoomServiceImpl) SearchRoomByAddress(address string) *responses.ResponseData {
	rooms, err := r.repo.SearchRoomByAddress(context.Background(), &address)

	if len(rooms) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    responses.StatusNoData,
			Data:       []dataaccess.SearchRoomByAddressRow{},
		}
	}
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
		Data: map[string]interface{}{
			"count": len(rooms),
			"rooms": rooms,
		},
	}
}

type BankAPI struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	Bin               string `json:"bin"`
	ShortName         string `json:"shortName"`
	Logo              string `json:"logo"`
	TransferSupported int    `json:"transferSupported"`
	LookupSupported   int    `json:"lookupSupported"`
	SwiftCode         string `json:"swift_code"`
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

// CheckUserLikedRoom implements RoomService.
func (r *RoomServiceImpl) CheckUserLikedRoom(roomId int, userId int) *responses.ResponseData {
	param := dataaccess.CheckUserLikedRoomParams{
		RoomID: int32(roomId),
		UserID: int32(userId),
	}
	_, err := r.repo.CheckUserLikedRoom(context.Background(), param)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    responses.StatusSuccess,
			Data:       false,
		}
	} else {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    responses.StatusSuccess,
			Data:       true,
		}
	}
}
