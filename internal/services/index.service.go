package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)


type IndexService interface{
	GetAllIndex(userid int32, month int32, year int32, mType string) *responses.ResponseData
	CreateIndex(userid int32,body *dataaccess.UpsertIndexParams) *responses.ResponseData
}