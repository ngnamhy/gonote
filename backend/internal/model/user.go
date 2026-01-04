package model

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type User struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username   string    `json:"username" gorm:"size:50;uniqueIndex;not null"`
	Email      string    `json:"email" gorm:"size:100;uniqueIndex;not null"`
	Password   string    `json:"-" gorm:"size:255;not null"`
	IsDisabled bool      `json:"is_disabled" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Role       Role      `json:"role" gorm:"size:20;default:user;not null,check:role in ('user','admin')"`
	AvatarURL  string    `json:"avatar_url" gorm:"size:255"`
}
