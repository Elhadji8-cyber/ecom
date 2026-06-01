package handler

import (
	"net/http"
	"server/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	authService service.AuthService
}

func NewRefreshTokenHandler(authService service.AuthService) *RefreshTokenHandler {
	return &RefreshTokenHandler{authService: authService}
}

func (h *RefreshTokenHandler) Handle(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
