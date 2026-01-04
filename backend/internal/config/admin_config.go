package config

import "os"

type AdminConfig struct {
	Email    string
	Username string
	Password string
}

func NewAdminConfig() *AdminConfig {
	return &AdminConfig{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}
}
