package service

import (
	"gonote/internal/model"
	"gonote/internal/repository"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	user_repo *repository.UserRepository
}

func NewUserService(user_repo *repository.UserRepository) *UserService {
	return &UserService{
		user_repo: user_repo,
	}
}

func (us *UserService) GetAllUser() ([]model.User, error) {
	users, error := us.user_repo.GetAllUser()
	return users, error
}

func (us *UserService) GetByID(id uuid.UUID) (model.User, error) {
	return us.user_repo.GetByID(id)
}

func (us *UserService) Create(username, password, email string) (*model.User, error) {
	id := uuid.New()
	createdAt := time.Now()
	user := &model.User{
		ID:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: createdAt,
	}
	return user, us.user_repo.Create(user)
}
