package model

import (
	"time"

	"github.com/google/uuid"
)

type VoteTarget string

const (
	VoteTargetPost    VoteTarget = "post"
	VoteTargetComment VoteTarget = "comment"
)

type Vote struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index"`
	TargetType VoteTarget `gorm:"type:varchar(20);not null;check:target_type IN ('post','comment')"`
	TargetID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	Value      int        `gorm:"not null;check:value IN (1,-1)"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
}
