package middleware

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateMiddlewareConfig() echo.MiddlewareFunc {
	signingKey := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

	config := middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("method %v tidak sesuai", token.Header["alg"])
				}

				return signingKey, nil
			})
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("JWT token tidak valid")
			}
			return token, nil
		},
	}

	return middleware.JWTWithConfig(config)
}
