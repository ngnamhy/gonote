package handler

import (
	"gonote/internal/dto"
	"gonote/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	user_service *service.UserService
}

func NewUserHandler(user_service *service.UserService) *UserHandler {
	return &UserHandler{user_service: user_service}
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.user_service.GetAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	type GetByIDParams struct {
		ID string `uri:"id" binding:"uuid"`
	}
	var params GetByIDParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, _ := uuid.Parse(params.ID)

	user, err := h.user_service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *UserHandler) Create(c *gin.Context) {
	var params dto.CreateUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.user_service.Create(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var params dto.UpdateUserParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	user, err := h.user_service.Update(id, params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid update",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"id":          id,
		"username":    user.Username,
		"email":       user.Email,
		"is_disabled": user.IsDisabled,
	})
}
