package service

import (
	"gonote/internal/dto"
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

func (us *UserService) Create(params dto.CreateUserParams) (*model.User, error) {
	id := uuid.New()
	createdAt := time.Now()
	user := &model.User{
		ID:        id,
		Username:  *params.Username,
		Password:  *params.Password,
		Email:     *params.Email,
		IsDisabled: false,
		CreatedAt: createdAt,
	}
	return user, us.user_repo.Create(user)
}

func (us *UserService) Update(id uuid.UUID, params dto.UpdateUserParams) (*model.User, error) {
	user, err := us.user_repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if params.Username != nil {
		user.Username = *params.Username
	}

	if params.Password != nil {
		user.Password = *params.Password
	}

	if params.Email != nil {
		user.Email = *params.Email
	}

	if params.IsDisabled != nil {
		user.IsDisabled = *params.IsDisabled
	}

	return &user, us.user_repo.Update(&user)
}
