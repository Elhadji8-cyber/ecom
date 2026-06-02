package handler

import (
	"net/http"
	"server/internal/auth/dto"
	"server/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	authService service.AuthService
}

func NewRegisterHandler(authService service.AuthService) *RegisterHandler {
	return &RegisterHandler{authService: authService}
}

func (h *RegisterHandler) Handle(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.authService.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}
