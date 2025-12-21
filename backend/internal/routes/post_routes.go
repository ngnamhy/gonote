package routes

import (
	"gonote/internal/handler"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	post_handler *handler.PostHandler
}

func NewPostRoutes(post_handler *handler.PostHandler) *PostRoutes {
	return &PostRoutes{
		post_handler: post_handler,
	}
}

func (pr *PostRoutes) Register(r *gin.RouterGroup) {
	posts := r.Group("/posts")

	h := pr.post_handler
	{
		posts.GET("", h.GetAllPost)
		posts.GET("/:id", h.GetByID)
		posts.POST("", h.Create)
		posts.PUT("/:id", h.Update)
		// users.DELETE("/:id", h.Delete)
	}
}
