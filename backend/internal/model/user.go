package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username   string    `gorm:"size:50;uniqueIndex;not null"`
	Email      string    `gorm:"size:100;uniqueIndex;not null"`
	Password   string    `gorm:"size:255;not null"`
	IsDisabled bool      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
