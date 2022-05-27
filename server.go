package main

import (
	"log"

	"github.com/YogiPristiawan/go-todo-api/infrastructures/databases/mysql"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/http"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/middleware"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/security/tokenize"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/validator"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/validator/translate"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// conect to database
	db := mysql.CreateConnection()

	// create server instance
	server := http.CreateServer()

	// create validator and validator translation instance
	validator := validator.CreateRequestValidator()
	trans := translate.CreateRequestValidatorTranslate(validator)

	// create tokenize
	tokenize := tokenize.NewJwtToken()

	// create middleware config
	jwtMiddleware := middleware.CreateMiddlewareConfig()

	// register routes
	api.CreateRoutes(server, db, validator, trans, jwtMiddleware, tokenize)
	server.Logger.Fatal(server.Start(":8080"))
}
