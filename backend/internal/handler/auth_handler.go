package handler

import (
	"gonote/internal/dto"
	"gonote/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

func NewAuthHandler(authService *service.AuthService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{authService: authService,
		userService: userService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var params dto.LoginParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	accessToken, refreshToken, user, err := h.authService.Login(&params)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong username or password"})
		return
	}

	c.SetCookie(
		"refresh_token",
		*refreshToken,
		7*24*60*60,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message":      "login success",
		"access_token": accessToken,
		"user":         user,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var params dto.RegisterParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	if *params.Password != *params.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password not match",
		})
		return
	}

	accessToken, refreshToken, user, err := h.authService.Register(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.SetCookie(
		"refresh_token",
		*refreshToken,
		7*24*60*60,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusCreated, gin.H{
		"message":      "register success",
		"access_token": accessToken,
		"user":         user,
	})

}

func (h *AuthHandler) Logout(c *gin.Context) {
}
