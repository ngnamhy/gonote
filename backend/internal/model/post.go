package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Title     string    `json:"title" gorm:"type:text;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Body      *string   `json:"body,omitempty" gorm:"type:text"`
	Score     int       `json:"score" gorm:"default:0"`
}
