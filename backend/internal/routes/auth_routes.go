package routes

import (
	"gonote/internal/handler"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	handler *handler.AuthHandler
}

func NewAuthRoutes(handler *handler.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		handler: handler,
	}
}

func (h *AuthRoutes) Register(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("login", h.handler.Login)
		auth.POST("logout", h.handler.Logout)
		auth.POST("register", h.handler.Register)
	}
}
