package routes

import (
	"server/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	Register       *handler.RegisterHandler
	Login          *handler.LoginHandler
	ForgotPassword *handler.ForgotPasswordHandler
	RefreshToken   *handler.RefreshTokenHandler
}

func RegisterAuthRoutes(router *gin.Engine, handlers AuthHandlers) {
	authGroup := router.Group("/auth")

	{
		authGroup.POST("/register", handlers.Register.Handle)
		authGroup.POST("/login", handlers.Login.Handle)
		authGroup.POST("/forgot-password", handlers.ForgotPassword.Handle)
		authGroup.POST("/refresh-token", handlers.RefreshToken.Handle)
	}
}
