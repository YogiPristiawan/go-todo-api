package profile

import (
	"strings"

	"github.com/YogiPristiawan/go-todo-api/applications/helpers"
	"github.com/YogiPristiawan/go-todo-api/domains/profile"
	tokenizeInfra "github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	useCase  profile.ProfileUseCase
	tokenize *tokenizeInfra.JwtToken
}

func (p *profileHandler) getProfile(c echo.Context) error {
	authorization := c.Request().Header["Authorization"]
	token := strings.Split(authorization[0], " ")[1]

	claims, _ := p.tokenize.DecodeAccessToken(token)

	result, err := p.useCase.GetProfile(claims.UserId)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return helpers.ResponseJsonHttpOk(c, "user profile", result)
}
