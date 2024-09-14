package common

import (
	"errors"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (*responses.UserResponse, error) {

	user, ok := c.Get("currentUser")
	if !ok {
		
		return nil, errors.New("Unauthorized")
	}

	currentUser, ok := user.(responses.UserResponse)
	if !ok {
		return nil, errors.New("user type assertion failed")
	}
	return &currentUser, nil

}