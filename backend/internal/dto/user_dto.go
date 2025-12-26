package dto

type CreateUserParams struct {
	Username        *string `json:"username" binding:"min=3,max=256,required"`
	Email           *string `json:"email" binding:"required,email"`
	Password        *string `json:"password" binding:"min=8,max=64,required"`
	ConfirmPassword *string `json:"confirm_password" binding:"min=8,max=64,required"`
}

type UpdateUserParams struct {
	Username   *string `json:"username" binding:"omitempty,min=3,max=256"`
	Email      *string `json:"email" binding:"omitempty,email"`
	Password   *string `json:"password" binding:"omitempty,min=8,max=64"`
	IsDisabled *bool   `json:"is_disabled" binding:"omitempty,boolean"`
}
