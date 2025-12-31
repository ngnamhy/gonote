package config

import (
	"context"
	mylog "gonote/pkg/log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

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
	RedisConfig   RedisConfig
}

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisClient(cfg RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Username:     cfg.Username,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     20,
		MinIdleConns: 5,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		mylog.AppLogger.Fatal().Msg("Error connecting to redis")
	}

	mylog.AppLogger.Info().Msg("Connected Redis")

	return client
}

func NewConfig() *Config {
	DB := DB{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	redisDB_str := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDB_str)
	if err != nil {
		redisDB = 0
	}
	RedisConfig := RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	}

	return &Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		LogDir:        os.Getenv("LOGDIR"),
		DB:            DB,
		RedisConfig:   RedisConfig,
	}
}
