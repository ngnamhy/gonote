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

	DB, err := db.New(&cfg.DB)

	if err != nil {
		mylog.AppLogger.Fatal().Msg("Error connecting to database")
	}

	appContext := app.NewAppContext(DB)

	a := app.NewApplication(cfg, appContext)

	if err := a.Run(); err != nil {
		mylog.AppLogger.Fatal().Err(err)
	}
}
