package mylog

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger

type LoggerConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func InitLogger(config LoggerConfig) {
	Logger = NewLogger(config)
}

func NewLogger(config LoggerConfig) *zerolog.Logger {
	lvl, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)

	writer := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}

	logger := zerolog.New(writer).With().Timestamp().Logger()

	return &logger
}
