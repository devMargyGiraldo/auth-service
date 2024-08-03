package main

import (
	"auth_service/internal/infrastructure"
	"auth_service/internal/interfaces/handler"
	"auth_service/internal/interfaces/repository"
	"auth_service/internal/usecase"
)

func main() {
	// Setup repository and use case
	userRepo := repository.NewInMemoryUserRepo()
	authUseCase := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)

	// Setup router
	router := infrastructure.NewRouter(authHandler)

	// Start the server
	infrastructure.StartServer(router)
}
