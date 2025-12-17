package service

import "gonote/internal/repository"

type UserService struct {
	user_repo *repository.UserRepository
}

func NewUserService(user_repo *repository.UserRepository) *UserService {
	return &UserService{
		user_repo: user_repo,
	}
}
