package dto

import "github.com/google/uuid"

type CreatePostParams struct {
	UserID uuid.UUID `json:"user_id" binding:"required,uuid"`
	Title  string    `json:"title" binding:"required,min=1,max=256"`
	Body   *string   `json:"body" binding:"omitempty"`
}

type UpdatePostParams struct {
	Title *string `json:"title" binding:"omitempty,min=1,max=256"`
	Body  *string `json:"body" binding:"omitempty"`
	Score *int    `json:"score" binding:"omitempty"`
}
