package config

import (
	"context"
	mylog "gonote/pkg/log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func NewRedisConfig() *RedisConfig {
	redisDB_str := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDB_str)
	if err != nil {
		redisDB = 0
	}
	return &RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	}
}

func (cfg *RedisConfig) NewRedisClient() *redis.Client {
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
