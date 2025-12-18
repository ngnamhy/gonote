package model

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `gorm:"size:50;uniqueIndex;not null"`
}
