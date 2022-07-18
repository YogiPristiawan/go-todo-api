package helper

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type customClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userId uint) string {
	ACCESS_TOKEN_EXPIRE_TIME, _ := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRE_TIME"), 10, 64)
	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")

	now := time.Now()
	ttl := time.Duration(ACCESS_TOKEN_EXPIRE_TIME) * time.Second

	claims := &customClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(ttl).Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(ACCESS_TOKEN_SECRET))

	return ss
}

func DecodeAccessToken(tokenString string) (*customClaims, error) {
	ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ACCESS_TOKEN_SECRET), nil
	})
	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func DecodeAuthJwtPayload(c echo.Context) (claims *customClaims, err error) {
	authorizationHeader := c.Request().Header.Get("Authorization")
	token := strings.Split(authorizationHeader, " ")[1]

	claims, err = DecodeAccessToken(token)

	if err != nil {
		err = exceptions.NewAuthenticationError(err.Error())
		return
	}
	return
}
