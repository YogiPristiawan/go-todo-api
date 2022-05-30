package infrastructures

import (
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/encrypt"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Http *echo.Echo
}

func (s *Server) GetHttp() *echo.Echo {
	return s.Http
}

type Database struct {
	MySql *gorm.DB
}

func (d *Database) GetMysql() *gorm.DB {
	return d.MySql
}

type Validator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

func (v *Validator) GetValidator() *validator.Validate {
	return v.Validator
}

func (v *Validator) GetTranslator() ut.Translator {
	return v.Trans
}

type Security struct {
	Jwt          *tokenize.JwtToken
	HashPassword *encrypt.HashPassword
}

func (s *Security) GetJwt() *tokenize.JwtToken {
	return s.Jwt
}

func (s *Security) GetHashPassword() *encrypt.HashPassword {
	return s.HashPassword
}

type Middleware struct {
	Auth echo.MiddlewareFunc
}

func (m *Middleware) GetAuth() echo.MiddlewareFunc {
	return m.Auth
}
