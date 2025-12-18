package dto

type CreateUserParams struct {
	Username *string `json:"username" binding:"min=3,max=256,required"`
	Email    *string `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"min=8,max=64,required"`
}
