package main

import (
	"github.com/YogiPristiawan/go-todo-api/infrastructures/http"
	"github.com/YogiPristiawan/go-todo-api/interfaces/http/api"
)

func main() {
	server := http.CreateServer()
	api.CreateRoutes(server)
	server.Logger.Fatal(server.Start(":8080"))
}
