package main

import (
	"log"

	"github.com/victorluisca/go-todo-app/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
