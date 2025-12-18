package model

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PostID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	ParentID  *uuid.UUID `gorm:"type:uuid;index"`
	Body      string     `gorm:"type:text;not null"`
	Score     int        `gorm:"default:0"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}
