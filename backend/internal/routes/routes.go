package routes

import "github.com/gin-gonic/gin"

type Routes interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Routes) {
	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}

}
