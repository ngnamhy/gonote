package handler

import (
	"gonote/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	user_service *service.UserService
}

func NewUserHandler(user_service *service.UserService) *UserHandler {
	return &UserHandler{user_service: user_service}
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Create(c *gin.Context) {
}

func (h *UserHandler) List(c *gin.Context) {
}
