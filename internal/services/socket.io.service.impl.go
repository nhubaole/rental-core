package services

import (
	"encoding/json"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type SocketIOServiceImpl struct {
	repo *dataaccess.Queries
}



func NewSocketIOServiceImpl() SocketIOService {
	return &SocketIOServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

// SendMessage implements SocketIOService.
func (s *SocketIOServiceImpl) SendMessage(req requests.MessageReq) *responses.ResponseData {
	res, err := common.SendPostRequest(global.Config.NodeServer.Url+"send-message", req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       true,
		}
	}
	var nodeRes responses.NodeResponse
	if err := json.Unmarshal([]byte(res), &nodeRes); err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to parse response from Node.js",
			Data:       nil,
		}
	}
	if !nodeRes.Success {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    nodeRes.Message,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       res,
	}
}