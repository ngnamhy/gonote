package app

import (
	"gonote/internal/config"
	"gonote/internal/model"
	"gonote/internal/routes"
	"gonote/pkg/auth"
	"gonote/pkg/cache"
	mylog "gonote/pkg/log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		mylog.AppLogger.Err(err)
	}

	userModel := NewUserModule(appContext)
	postModel := NewPostModule(appContext)
	authModel := NewAuthModule(appContext)
	modules := []Module{
		userModel,
		postModel,
		authModel,
	}

	jwtService := auth.NewJWTService(appContext.redisCache)
	routes.RegisterRoutes(r, jwtService, GetRoutesFromModules(modules)...)
	seedData(config.AdminConfig, appContext.DB)

	return &Application{
		config: config,
		router: r,
	}
}

func seedData(adminConfig *config.AdminConfig, db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Where("role = ?", "admin").Count(&count)

	if count > 0 {
		mylog.AppLogger.Info().Msg("Admin user already exists, skipping seed.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminConfig.Password), bcrypt.DefaultCost)
	if err != nil {
		mylog.AppLogger.Panic().Err(err).Msg("Failed to hash admin password:")
	}

	adminUser := model.User{
		Username:   adminConfig.Username,
		Email:      adminConfig.Email,
		Password:   string(hashedPassword),
		Role:       "admin",
		IsDisabled: false,
		CreatedAt:  time.Now(),
	}

	if err := db.Create(&adminUser).Error; err != nil {
		mylog.AppLogger.Panic().Err(err).Msg("Failed to create admin user:")
	}

	mylog.AppLogger.Info().Msg("Admin user created successfully")
}

func (app *Application) Run() error {
	return app.router.Run(app.config.ServerAddress)
}
