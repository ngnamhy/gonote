package app

import (
	"gonote/internal/handler"
	"gonote/internal/repository"
	"gonote/internal/routes"
	"gonote/internal/service"
	"gonote/pkg/auth"
)

type AuthModule struct {
	authRoutes *routes.AuthRoutes
}

func NewAuthModule(ctx *AppContext) *AuthModule {
	userRepo := repository.NewUserRepository(ctx.DB)
	jwtService := auth.NewJWTService(ctx.redisCache)
	authService := service.NewAuthService(userRepo, jwtService, ctx.redisCache)
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(authService, userService)
	authRoutes := routes.NewAuthRoutes(authHandler)
	return &AuthModule{authRoutes: authRoutes}
}

func (am *AuthModule) Routes() routes.Routes {
	return am.authRoutes
}
