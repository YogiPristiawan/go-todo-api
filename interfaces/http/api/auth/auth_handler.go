package auth

import (
	"github.com/labstack/echo/v4"

	"github.com/YogiPristiawan/go-todo-api/applications/exceptions"
	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	"github.com/YogiPristiawan/go-todo-api/domains/auth/entities"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	useCase              auth.AuthUseCase
	validator            *validator.Validate
	validatorTranslation ut.Translator
}

func (a *authHandler) login(c echo.Context) error {
	// collect payload
	l := new(entities.AuthLoginRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, l); err != nil {
		return helpers.HandleError(c, err)
	}

	// validate payload
	if err := a.validator.Struct(l); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(a.validatorTranslation)

			for _, val := range errors {
				return helpers.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	payload := &entities.AuthLoginRequest{
		Username: l.Username,
		Password: l.Password,
	}

	result, err := a.useCase.Login(payload)
	if err != nil {
		return helpers.HandleError(c, exceptions.NewInvariantError(err.Error()))
	}

	return helpers.ResponseJsonHttpOk(c, "login success", result)
}

func (a *authHandler) register(c echo.Context) error {
	// collect payload
	r := new(entities.AuthRegisterRequest)
	if err := (&echo.DefaultBinder{}).BindBody(c, r); err != nil {
		return helpers.HandleError(c, err)
	}

	// validate payload
	if err := a.validator.Struct(r); err != nil {
		if he, ok := err.(validator.ValidationErrors); ok {
			errors := he.Translate(a.validatorTranslation)

			for _, val := range errors {
				return helpers.HandleError(c, exceptions.NewInvariantError(val))
			}
		}
	}

	payload := &entities.AuthRegisterRequest{
		Username:  r.Username,
		Password:  r.Password,
		Gender:    r.Gender,
		BirthDate: r.BirthDate,
	}

	result, err := a.useCase.Register(payload)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonCreated(c, "register success", result)

}
