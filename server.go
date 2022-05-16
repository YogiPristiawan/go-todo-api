package main

import (
	"log"

	"github.com/YogiPristiawan/go-todo-api/infrastructures/databases/mysql"
	"github.com/YogiPristiawan/go-todo-api/infrastructures/http"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db := mysql.CreateConnection()

	server := http.CreateServer()
	api.CreateRoutes(server, db)
	server.Logger.Fatal(server.Start(":8080"))
}
