package profile

import (
	"github.com/YogiPristiawan/go-todo-api/domains"
	profileDomain "github.com/YogiPristiawan/go-todo-api/domains/profile"
	"github.com/golobby/container/v3"
)

var profileUseCase profileDomain.ProfileUseCase
var security domains.Security
var middleware domains.Middleware
var server domains.Server

func InitRoutes() {
	container.Resolve(&profileUseCase)
	container.Resolve(&security)
	container.Resolve(&middleware)
	container.Resolve(&server)

	handler := &profileHandler{
		useCase:  profileUseCase,
		tokenize: security.GetJwt(),
	}

	e := server.GetHttp()
	authMiddleware := middleware.GetAuth()

	g := e.Group("/profile", authMiddleware)

	g.GET("", handler.getProfile)
}
