package app

import (
	"gonote/internal/handler"
	"gonote/internal/repository"
	"gonote/internal/routes"
	"gonote/internal/service"
)

type UserModule struct {
	userRoutes *routes.UserRoutes
}

func NewUserModule(appContext *AppContext) *UserModule {
	userRepo := repository.NewUserRepository(appContext.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)
	return &UserModule{
		userRoutes: userRoutes,
	}
}

func (um *UserModule) Routes() routes.Routes {
	return um.userRoutes
}
