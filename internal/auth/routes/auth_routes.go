package routes

import (
	"server/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
}
