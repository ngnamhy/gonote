package config

import (
	"os"
)

type Config struct {
	ServerAddress string
	JWTSecret     string
	LogDir        string
	DBConfig      *DBConfig
	RedisConfig   *RedisConfig
	AdminConfig   *AdminConfig
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		LogDir:        os.Getenv("LOGDIR"),
		DBConfig:      NewDBConfig(),
		RedisConfig:   NewRedisConfig(),
		AdminConfig:   NewAdminConfig(),
	}
}
