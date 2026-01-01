package app

import (
	"gonote/internal/config"
	"gonote/internal/model"
	"gonote/internal/routes"
	"gonote/pkg/auth"
	"gonote/pkg/cache"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB         *gorm.DB
	redisCache *cache.RedisCache
}

func NewAppContext(DB *gorm.DB, rdb *cache.RedisCache) *AppContext {
	return &AppContext{
		DB:         DB,
		redisCache: rdb,
	}
}

type Application struct {
	context *AppContext
	config  *config.Config
	router  *gin.Engine
}

func NewApplication(config *config.Config, appContext *AppContext) *Application {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	if err := appContext.DB.AutoMigrate(
		&model.User{},
		&model.Tag{},
		&model.Post{},
		&model.PostTag{},
		&model.Comment{},
		&model.Vote{},
	); err != nil {
		log.Panic(err)
	}

	modules := []Module{
		NewUserModule(appContext),
		NewPostModule(appContext),
		NewAuthModule(appContext),
	}

	jwtService := auth.NewJWTService(appContext.redisCache)
	routes.RegisterRoutes(r, jwtService, GetRoutesFromModules(modules)...)

	return &Application{
		config: config,
		router: r,
	}
}

func (app *Application) Run() error {
	return app.router.Run(app.config.ServerAddress)
}
