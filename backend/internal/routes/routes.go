package routes

import (
	"gonote/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Routes) {
	r.Use(middleware.LoggerMiddleware())
	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}

}
