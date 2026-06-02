package main

import (
	"log"
	"server/internal/auth/handler"
	"server/internal/auth/repository"
	"server/internal/auth/routes"
	"server/internal/auth/service"
	"server/internal/auth/validator"
	"server/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Initialize Database
	config.ConnectDatabase()

	// Initialize Services
	pwdSvc := service.NewPasswordService()
	jwtSvc := service.NewJwtService()
	authVal := validator.NewAuthValidator()

	// Initialize Repository
	authRepo := repository.NewAuthRepository(config.DB)

	// Initialize Service
	authSvc := service.NewAuthService(authRepo, pwdSvc, jwtSvc, authVal)

	// Initialize Handlers
	handlers := routes.AuthHandlers{
		Register:       handler.NewRegisterHandler(authSvc),
		Login:          handler.NewLoginHandler(authSvc),
		ForgotPassword: handler.NewForgotPasswordHandler(authSvc),
		RefreshToken:   handler.NewRefreshTokenHandler(authSvc),
	}

	// Initialize Router
	router := gin.Default()

	// Register Routes
	routes.RegisterAuthRoutes(router, handlers)

	// Start Server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
