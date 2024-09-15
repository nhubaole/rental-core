package common

import (
	"fmt"
	"os"
	"smart-rental/global"
	"smart-rental/pkg/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user responses.UserResponse) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"full_name": user.FullName,
		"sub":       user.PhoneNumber,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"authorize": true,
	})
	return token.SignedString([]byte(global.Config.JWT.SecretKey))
}

func TokenValid(ctx *gin.Context) error {
	tokenString := ctx.GetHeader("Authorization")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}
