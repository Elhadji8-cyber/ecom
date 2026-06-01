package handler

import (
	"net/http"
	"server/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	authService service.AuthService
}

func NewForgotPasswordHandler(authService service.AuthService) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{authService: authService}
}

func (h *ForgotPasswordHandler) Handle(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.ForgotPassword(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reset password email sent"})
}
