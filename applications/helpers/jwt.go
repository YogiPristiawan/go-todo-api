package helpers

import (
	"time"

	"github.com/YogiPristiawan/go-todo-api/applications/constants"
	"github.com/golang-jwt/jwt"
)

type customClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userId uint) string {
	claims := &customClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: constants.ACCCESS_TOKEN_EXPIRE_TIME,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, _ := token.SignedString([]byte(constants.ACCESS_TOKEN_SECRET))

	return ss
}
