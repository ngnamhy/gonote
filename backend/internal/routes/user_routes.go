package routes

import (
	"gonote/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	user_handler *handler.UserHandler
}

func NewUserRoutes(user_handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		user_handler: user_handler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")

	h := ur.user_handler
	{
		users.GET("", h.GetAllUser)
		users.GET("/:id", h.GetByID)
		users.POST("", h.Create)
		users.PUT("/:id", h.Update)
		// users.DELETE("/:id", h.Delete)
	}
}
