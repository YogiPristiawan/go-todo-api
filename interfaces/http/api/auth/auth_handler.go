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

type AuthHandler struct {
	useCase              auth.AuthUseCase
	validator            *validator.Validate
	validatorTranslation ut.Translator
}

func NewAuthHandler(
	useCase auth.AuthUseCase,
	validator *validator.Validate,
	validationTranslation ut.Translator,
) *AuthHandler {
	return &AuthHandler{
		useCase:              useCase,
		validator:            validator,
		validatorTranslation: validationTranslation,
	}
}

func (a *AuthHandler) Login(c echo.Context) error {
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
