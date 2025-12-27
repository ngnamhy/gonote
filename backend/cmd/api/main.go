package main

import (
	"gonote/internal/app"
	"gonote/internal/config"
	"gonote/internal/db"
	mylog "gonote/pkg/log"
	"path/filepath"
)

func main() {
	cfg := config.NewConfig()

	logFile := filepath.Join(cfg.LogDir, "app.log")

	mylog.InitLogger(mylog.LoggerConfig{
		Level:      "info",
		Filename:   logFile,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     5,
		Compress:   true,
	})

	mylog.Logger.Info().Msg("Hello World")

	DB, err := db.New(&cfg.DB)

	if err != nil {
		mylog.Logger.Fatal().Msg("Error connecting to database")
	}

	appContext := app.NewAppContext(DB)

	a := app.NewApplication(cfg, appContext)

	if err := a.Run(); err != nil {
		mylog.Logger.Fatal().Err(err)
	}
}
