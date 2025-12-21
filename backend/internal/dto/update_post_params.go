package dto

type UpdatePostParams struct {
	Title *string `json:"title" binding:"omitempty,min=1,max=256"`
	Body  *string `json:"body" binding:"omitempty"`
	Score *int    `json:"score" binding:"omitempty"`
}
