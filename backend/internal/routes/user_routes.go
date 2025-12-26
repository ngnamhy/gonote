package routes

import (
	"gonote/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler *handler.UserHandler
}

func NewUserRoutes(userHandler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		userHandler: userHandler,
	}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	users := r.Group("/users")

	h := ur.userHandler
	{
		users.GET("", h.GetAllUser)
		users.GET("/:id", h.GetByID)
		users.POST("", h.Create)
		users.PUT("/:id", h.Update)
		// users.DELETE("/:id", h.Delete)
	}
}
