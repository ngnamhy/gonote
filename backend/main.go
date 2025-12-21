package main

import (
	"gonote/internal/config"
	"gonote/internal/db"
	"gonote/internal/handler"
	"gonote/internal/repository"
	"gonote/internal/routes"
	"gonote/internal/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.NewConfig()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	DB, err := db.New()
	if err != nil {
		log.Println("Cannot connect to db, exitting")
		return
	}

	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := routes.NewUserRoutes(userHandler)

	postRepo := repository.NewPostRepository(DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)
	postRoutes := routes.NewPostRoutes(postHandler)

	routes.RegisterRoutes(r, userRoutes, postRoutes)

	r.Run(cfg.ServerAddress)
}
