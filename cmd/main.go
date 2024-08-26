package main

import (
	"github.com/shirocola/user-service/internal/infrastructure"
	"github.com/shirocola/user-service/internal/infrastructure/persistence"
	"github.com/shirocola/user-service/internal/interface/http"
	"github.com/shirocola/user-service/internal/usecase"
)

func main() {
	userRepo := persistence.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	// Initialize the router
	router := infrastructure.InitRouter(userHandler)

	// Start the server
	infrastructure.StartServer(router)
}
