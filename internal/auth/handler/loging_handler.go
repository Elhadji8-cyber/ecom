package handler

import (
	"net/http"
	"server/internal/auth/dto"
	"server/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	authService service.AuthService
}

func NewLoginHandler(authService service.AuthService) *LoginHandler {
	return &LoginHandler{authService: authService}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
