package model

import "time"

import "github.com/google/uuid"


type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
