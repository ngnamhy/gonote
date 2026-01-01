package routes

import (
	"gonote/internal/middleware"
	"gonote/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, jwtService *auth.JWTService, routes ...Routes) {
	api := r.Group("/api/v1")
	api.Use(middleware.LoggerMiddleware())

	protected := api.Group("")
	middleware.InitAuthMiddleware(jwtService)
	protected.Use(middleware.AuthMiddleware())

	for _, route := range routes {
		switch route.(type) {
		case *AuthRoutes:
			route.Register(api)
		default:
			route.Register(protected)
		}
	}

}
