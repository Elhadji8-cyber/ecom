package main

import (
	"log"
	authHandler "server/internal/auth/handler"
	authRepo "server/internal/auth/repository"
	authRoutes "server/internal/auth/routes"
	authService "server/internal/auth/service"
	authValidator "server/internal/auth/validator"
	"server/internal/config"
	productHandler "server/internal/product/handler"
	productRepo "server/internal/product/repository"
	productRoutes "server/internal/product/routes"
	productService "server/internal/product/service"

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

	// Initialize Auth Components
	authPwdSvc := authService.NewPasswordService()
	authJwtSvc := authService.NewJwtService()
	authVal := authValidator.NewAuthValidator()
	aRepo := authRepo.NewAuthRepository(config.DB)
	aService := authService.NewAuthService(aRepo, authPwdSvc, authJwtSvc, authVal)
	aHandlers := authRoutes.AuthHandlers{
		Register:       authHandler.NewRegisterHandler(aService),
		Login:          authHandler.NewLoginHandler(aService),
		ForgotPassword: authHandler.NewForgotPasswordHandler(aService),
		RefreshToken:   authHandler.NewRefreshTokenHandler(aService),
	}

	// Initialize Product Components
	pRepo := productRepo.NewProductRepository(config.DB)
	pService := productService.NewProductService(pRepo)
	pHandler := productHandler.NewProductHandler(pService)

	// Initialize Router
	router := gin.Default()

	// Register Routes
	authRoutes.RegisterAuthRoutes(router, aHandlers)
	productRoutes.RegisterProductRoutes(router, pHandler)

	// Start Server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
