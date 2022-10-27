package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type customClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateAccessToken(userId int64) string {
	var ACCESS_TOKEN_EXPIRE_TIME, _ = strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRE_TIME"), 10, 64)
	var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")

	now := time.Now()
	ttl := time.Duration(ACCESS_TOKEN_EXPIRE_TIME) * time.Second

	claims := customClaims{
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

func DecodeAccessToken(accessToken string) (customClaims, error) {
	var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
	token, err := jwt.ParseWithClaims(accessToken, customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ACCESS_TOKEN_SECRET), nil
	})
	if claims, ok := token.Claims.(customClaims); ok && token.Valid {
		return claims, nil
	} else {
		return customClaims{}, err
	}
}
