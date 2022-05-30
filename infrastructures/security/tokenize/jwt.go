package tokenize

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtToken struct {
}

type customClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func NewJwtToken() *JwtToken {
	return &JwtToken{}
}

func (j *JwtToken) GenerateAccessToken(userId uint) string {
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

func (t *JwtToken) DecodeAccessToken(tokenString string) (*customClaims, error) {
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

func (t *JwtToken) TestFunc() {
	fmt.Printf("Infrastructures.Tokenzie: Hello World")
}
