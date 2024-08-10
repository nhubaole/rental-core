package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}){
	c.JSON(http.StatusOK, ResponseData{
		StatusCode: code,
		Message: "Success",
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string){ 
	c.JSON(http.StatusOK, ResponseData{
		StatusCode: code,
		Message: message,
		Data: nil,
	})
}
func APIResponse(c *gin.Context, code int, msg string, data interface{}){
	c.JSON(code, ResponseData{
		StatusCode: code,
		Message: msg,
		Data: data,
	})
}