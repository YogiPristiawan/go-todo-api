package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/YogiPristiawan/go-todo-api/domain/auth"
	"github.com/YogiPristiawan/go-todo-api/modules/exceptions"
	"github.com/YogiPristiawan/go-todo-api/modules/helper"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	UseCase              auth.AuthUseCase
	Validator            *validator.Validate
	ValidatorTranslation ut.Translator
}

func (a *AuthHandler) Login(c echo.Context) error {
	// collect payload
	l := new(auth.LoginRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, l); err != nil {
		return helper.HandleError(c, err)
	}

	// validate payload
	if err := a.Validator.Struct(l); err != nil {
		err = helper.ValidatorErrorTranslate(err, a.ValidatorTranslation)
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := a.UseCase.Login(&auth.LoginRequest{
		Username: l.Username,
		Password: l.Password,
	})
	if err != nil {
		return helper.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	return helper.ResponseJsonHttpOk(c, "login success", result)
}

func (a *AuthHandler) Register(c echo.Context) error {
	// collect payload
	r := new(auth.RegisterRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helper.HandleError(c, err)
	}

	// validate payload
	if err := a.Validator.Struct(r); err != nil {
		err = helper.ValidatorErrorTranslate(err, a.ValidatorTranslation)
		return helper.HandleError(c, err)
	}

	// call use case
	result, err := a.UseCase.Register(&auth.RegisterRequest{
		Username:  r.Username,
		Password:  r.Password,
		Gender:    r.Gender,
		BirthDate: r.BirthDate,
	})
	if err != nil {
		return helper.HandleError(c, err)
	}

	return helper.ResponseJsonCreated(c, "register success", result)

}
