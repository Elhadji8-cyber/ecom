package routes

import (
	"server/internal/product/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, h *handler.ProductHandler) {
	productGroup := router.Group("/products")
	{
		productGroup.POST("/", h.Create)
		productGroup.GET("/", h.GetAll)
		productGroup.GET("/:id", h.GetByID)
		productGroup.PUT("/:id", h.Update)
		productGroup.DELETE("/:id", h.Delete)
		productGroup.POST("/categories", h.CreateCategory)
	}
}
