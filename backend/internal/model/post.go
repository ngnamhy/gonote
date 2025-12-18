package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Title     string    `gorm:"type:text;not null"`
	Body      *string   `gorm:"type:text"`
	Score     int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
