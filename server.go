package main

import (
	"log"

	"github.com/YogiPristiawan/go-todo-api/applications/use_case"
	"github.com/YogiPristiawan/go-todo-api/domains"
	"github.com/YogiPristiawan/go-todo-api/domains/auth"
	"github.com/YogiPristiawan/go-todo-api/domains/profile"
	"github.com/YogiPristiawan/go-todo-api/domains/todo"
	"github.com/YogiPristiawan/go-todo-api/domains/users"
	"github.com/YogiPristiawan/go-todo-api/infrastructures"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/databases/mysql"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/http"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/repository"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/encrypt"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/middleware"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/validator"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/validator/translate"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api"
	"github.com/golobby/container/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Container
	registerServer()
	registerDatabase()
	registerValidator()
	registerSecurity()
	registerRepository()
	registerUseCase()

	// register routes
	api.CreateRoutes()
	var server domains.Server
	container.Resolve(&server)
	server.GetHttp().Logger.Fatal(server.GetHttp().Start(":8080"))
}

func registerServer() {
	// Server
	container.Singleton(func() domains.Server {
		return &infrastructures.Server{
			Http: http.CreateServer(),
		}
	})
}

func registerDatabase() {
	// Database
	container.Singleton(func() domains.Database {
		return &infrastructures.Database{
			MySql: mysql.CreateConnection(),
		}
	})
}

func registerValidator() {
	// validator
	validator := validator.CreateRequestValidator()
	trans := translate.CreateRequestValidatorTranslate(validator)
	container.Singleton(func() domains.Validator {
		return &infrastructures.Validator{
			Validator: validator,
			Trans:     trans,
		}
	})
}

func registerSecurity() {
	// encrypt
	container.Singleton(func() domains.Security {
		return &infrastructures.Security{
			Jwt:          tokenize.NewJwtToken(),
			HashPassword: &encrypt.HashPassword{},
		}
	})
	// middleware
	container.Singleton(func() domains.Middleware {
		return &infrastructures.Middleware{
			Auth: middleware.CreateAuthMiddleware(),
		}
	})
}

func registerRepository() {
	// Repositories
	container.Singleton(func(db domains.Database) users.UserRepository {
		return &repository.UserRepository{
			DB: db.GetMysql(),
		}
	})
	container.Singleton(func(db domains.Database) todo.TodoRepository {
		return &repository.TodoRepository{
			DB: db.GetMysql(),
		}
	})
}

func registerUseCase() {
	// Use Case
	container.Singleton(func(r users.UserRepository, t domains.Security) users.UserUseCase {
		return &use_case.UserUseCase{
			UserRepository: r,
			Tokenize:       t.GetJwt(),
		}
	})
	container.Singleton(func(r users.UserRepository, t domains.Security) auth.AuthUseCase {
		return &use_case.AuthUseCase{
			UserRepository:  r,
			Tokenize:        t.GetJwt(),
			PasswordManager: t.GetHashPassword(),
		}
	})
	container.Singleton(func(r users.UserRepository) profile.ProfileUseCase {
		return &use_case.ProfileUseCase{
			UserRepository: r,
		}
	})
	container.Singleton(func(r todo.TodoRepository, t domains.Security) todo.TodoUseCase {
		return &use_case.TodoUseCase{
			TodoRepository: r,
		}
	})

}
