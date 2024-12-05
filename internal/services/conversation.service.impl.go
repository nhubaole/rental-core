package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ConversationServiceImpl struct {
	repo *dataaccess.Queries
}

func NewConversationServiceImpl() ConversationService {
	return &ConversationServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

// Create implements ConversationService.
func (c *ConversationServiceImpl) Create(req requests.CreateConversationReq, userID int) *responses.ResponseData {
	var params dataaccess.CreateConversationParams

	common.MapStruct(req, &params)
	params.UserA = int32(userID)
	id,err := c.repo.CreateConversation(context.Background(), params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       id,
	}
}

// GetByUserID implements ConversationService.
func (c *ConversationServiceImpl) GetByUserID(userID int) *responses.ResponseData {
	var res []responses.ConversationRes
	conversations, err := c.repo.GetConversationByUserID(context.Background(), int32(userID))
	for _, conversation := range conversations {
		
		lastMessage, messErr := c.repo.GetMessageByID(context.Background(), *conversation.LastMessageID)
		if messErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
		conversationRes := responses.ConversationRes {
			ID: conversation.ID,
			UserA: conversation.UserA,
			UserB: conversation.UserB,
			LastMessage: lastMessage,
		}

		res = append(res, conversationRes)
	}
	
	if len(conversations) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}
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
		Data:       res,
	}
}
