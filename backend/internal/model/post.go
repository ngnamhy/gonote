package model

import "time"

type Post struct {
	ID        uint64 `gorm:"primaryKey"`
	UserID    uint64 `gorm:"not null"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"type:text;not null"`
	ViewCount int
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
