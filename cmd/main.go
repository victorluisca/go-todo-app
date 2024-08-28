package main

import "github.com/victorluisca/go-todo-app/cmd/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
