package middlewares

import (
	"context"
	"fmt"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"

	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenMiddleware(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	jwtToken := strings.Split(tokenString, " ")

	if len(jwtToken) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	bearerToken, _ := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(global.Config.JWT.SecretKey), nil
	})
	// fmt.Print(bearerToken.Valid)

	if claims, ok := bearerToken.Claims.(jwt.MapClaims); ok && bearerToken.Valid {

		// Check the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		//var user dataaccess.User

		db := dataaccess.New(global.Db)
		currentUser, err := db.GetUserByPhone(context.Background(),claims["sub"].(string))
		if err != nil {
			panic(err)
		}

		if currentUser.PhoneNumber == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		ctx.Set("currentUser", currentUser)

		//Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}