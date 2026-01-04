package main

import (
	"gonote/internal/app"
	"gonote/internal/config"
	"gonote/internal/db"
	"gonote/pkg/cache"
	mylog "gonote/pkg/log"
	"path/filepath"
)

func main() {
	cfg := config.NewConfig()

	appLogFile := filepath.Join(cfg.LogDir, "app.log")
	httpLogFile := filepath.Join(cfg.LogDir, "http.log")

	mylog.InitLoggers(
		mylog.LoggerConfig{
			Level:      "info",
			Filename:   appLogFile,
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		},
		mylog.LoggerConfig{
			Level:      "info",
			Filename:   httpLogFile,
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     7,
			Compress:   true,
		},
	)

	DB, err := db.New(cfg.DBConfig)

	if err != nil {
		mylog.AppLogger.Fatal().Msg("Error connecting to database")
	}

	redisClient := cfg.RedisConfig.NewRedisClient()
	redisCache := cache.NewRedisCache(redisClient)

	appContext := app.NewAppContext(DB, redisCache)

	a := app.NewApplication(cfg, appContext)

	if err := a.Run(); err != nil {
		mylog.AppLogger.Fatal().Err(err)
	}
}
