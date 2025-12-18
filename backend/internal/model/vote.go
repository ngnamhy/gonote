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
	TargetType VoteTarget `gorm:"type:vote_target;not null"`
	TargetID   uuid.UUID  `gorm:"type:uuid;not null;index"`
	Value      int        `gorm:"not null;check:value IN (1,-1)"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
}
