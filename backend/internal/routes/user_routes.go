package routes

import "gonote/internal/handler"

type UserRoutes struct {
	user_handler *handler.UserHandler
}

func NewUserRoutes(user_handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		user_handler: user_handler,
	}
}
