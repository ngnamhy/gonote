package mylog

import (
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	AppLogger  *zerolog.Logger // Log chung của ứng dụng
	HTTPLogger *zerolog.Logger // Log riêng cho HTTP request
)

type LoggerConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func InitLoggers(appConfig, httpConfig LoggerConfig) {
	AppLogger = NewLogger(appConfig)
	HTTPLogger = NewLogger(httpConfig)
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

	logger := zerolog.New(writer).With().Caller().Timestamp().Logger()

	return &logger
}
