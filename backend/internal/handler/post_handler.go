package handler

import (
	"gonote/internal/dto"
	"gonote/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetAllPost(c *gin.Context) {
	posts, err := h.postService.GetAllPost()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetByID(c *gin.Context) {
	type GetByIDParams struct {
		ID string `uri:"id" binding:"uuid"`
	}
	var params GetByIDParams

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, _ := uuid.Parse(params.ID)

	post, err := h.postService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": post})
}

func (h *PostHandler) Create(c *gin.Context) {
	var params dto.CreatePostParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.postService.Create(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var params dto.UpdatePostParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	post, err := h.postService.Update(id, params)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid update",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"post": post,
	})

}
