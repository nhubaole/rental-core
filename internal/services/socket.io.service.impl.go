package services

import (
	"fmt"
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
	fmt.Print(res)
}