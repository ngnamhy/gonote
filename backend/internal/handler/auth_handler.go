package handler

import (
	"gonote/internal/dto"
	"gonote/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var params dto.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	accessToken, err := h.authService.Login(params)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "login success",
		"access_token": accessToken})
}

func (h *AuthHandler) Logout(c *gin.Context) {
}
