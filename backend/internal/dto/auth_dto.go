package dto

import (
	"gonote/internal/model"
	"time"
)

type LoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterParams struct {
	Username        *string `json:"username" binding:"min=3,max=256,required"`
	Email           *string `json:"email" binding:"required,email"`
	Password        *string `json:"password" binding:"min=8,max=64,required"`
	ConfirmPassword *string `json:"confirm_password" validate:"eqcsfield=Password"`
}

func ToCreateUserParams(registerParams *RegisterParams) *CreateUserParams {
	return &CreateUserParams{
		Username: registerParams.Username,
		Email:    registerParams.Email,
		Password: registerParams.Password,
	}
}

func ToUser(params *RegisterParams) *model.User {
	return &model.User{
		Username:   *params.Username,
		Email:      *params.Email,
		Password:   *params.Password,
		IsDisabled: false,
		AvatarURL:  "",
		CreatedAt:  time.Now(),
		Role:       "user",
	}
}
