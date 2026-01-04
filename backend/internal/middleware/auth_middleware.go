package middleware

import (
	"gonote/pkg/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	jwtService *auth.JWTService
)

func InitAuthMiddleware(jwtServiceInjected *auth.JWTService) {
	jwtService = jwtServiceInjected
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || strings.HasPrefix(authHeader, "Bearer ") == false {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "Authorization header is missing or invalid"},
			)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		jwtService.ParseAccessToken(token)
		_, _, err := jwtService.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"message": "Authorization header is missing or invalid"},
			)
			return
		}
		c.Next()
	}
}
