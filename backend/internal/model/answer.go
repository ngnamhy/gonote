package model

import "time"

type Answer struct {
	ID        uint64 `gorm:"primaryKey"`
	PostID    uint64 `gorm:"not null"`
	UserID    uint64 `gorm:"not null"`
	Body      string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
