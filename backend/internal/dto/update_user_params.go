package dto

type UpdateUserParams struct {
	Username   *string `json:"username" binding:"omitempty,min=3,max=256"`
	Email      *string `json:"email" binding:"omitempty,email"`
	Password   *string `json:"password" binding:"omitempty,min=8,max=64"`
	IsDisabled *bool   `json:"is_disabled" binding:"omitempty,boolean"`
}
