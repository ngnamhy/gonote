package config

import "os"

type DB struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type Config struct {
	ServerAddress string
	JWTSecret     string
	LogDir        string
	DB            DB
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		LogDir:        os.Getenv("LOGDIR"),
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
	}
}
