package responses

import (

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	StatusCode int         `json:"errCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func APIResponse(c *gin.Context, code int, msg string, data interface{}){
	c.JSON(code, ResponseData{
		StatusCode: code,
		Message: msg,
		Data: data,
	})
}