package domains

import (
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/encrypt"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server interface {
	GetHttp() *echo.Echo
}

type Database interface {
	GetMysql() *gorm.DB
}

type Validator interface {
	GetValidator() *validator.Validate
	GetTranslator() ut.Translator
}

type Security interface {
	GetJwt() *tokenize.JwtToken
	GetHashPassword() *encrypt.HashPassword
}

type Middleware interface {
	GetAuth() echo.MiddlewareFunc
}
