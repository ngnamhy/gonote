package main

import (
	"gonote/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.NewConfig()
	r := gin.Default()

	r.Run(cfg.ServerAddress)
}
