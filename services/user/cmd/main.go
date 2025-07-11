package main

import (
	"log"

	"github.com/focuscw0w/microservices/services/user/cmd/api"
	"github.com/focuscw0w/microservices/services/user/internal/service"
)

type application struct {
	userService *internal.UserService
}

func main() {
	server := api.NewAPIServer(":8080", nil)

	// create storage
	// create service
	// create handler

	err := server.Run() 
	if err != nil {
		log.Fatal("Failed to start server:", err)
		return
	}
	log.Println("Server is running on :8080")
}