package common

import (
	"errors"
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
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"authorize": true,
	})
	return token.SignedString([]byte(global.Config.JWT.SecretKey))
}

func TokenValid(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := ctx.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateLandlordRoleJWT(ctx *gin.Context) error {
	token, err := TokenValid(ctx)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 0 {
		return nil
	}
	return errors.New("invalid landlord token provided")
}

func ValidateTenantRoleJWT(ctx *gin.Context) error {
	token, err := TokenValid(ctx)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid tenant token provided")
}