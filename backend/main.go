package main

import (
	"gonote/internal/config"
	"gonote/internal/db"
	"gonote/internal/handler"
	"gonote/internal/repository"
	"gonote/internal/routes"
	"gonote/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.NewConfig()

	r := gin.Default()

	DB, err := db.New()
	if err != nil {
		panic(err)
	}

	user_repo := repository.NewUserRepository(DB)

	user_service := service.NewUserService(user_repo)

	user_handler := handler.NewUserHandler(user_service)

	user_routes := routes.NewUserRoutes(user_handler)

	routes.RegisterRoutes(r, user_routes)

	r.Run(cfg.ServerAddress)
}
