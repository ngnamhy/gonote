package service

import (
	"gonote/internal/dto"
	"gonote/internal/model"
	"gonote/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) GetAllUser() ([]model.User, error) {
	users, err := us.userRepo.GetAllUser()
	return users, err
}

func (us *UserService) GetByID(id *uuid.UUID) (model.User, error) {
	return us.userRepo.GetByID(id)
}

func (us *UserService) GetByUsername(username *string) (*model.User, error) {
	return us.userRepo.GetByUsername(username)
}

func (us *UserService) Create(params *dto.CreateUserParams) (*model.User, error) {
	id := uuid.New()
	createdAt := time.Now()

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(*params.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:         id,
		Username:   *params.Username,
		Password:   string(hashed),
		Email:      *params.Email,
		IsDisabled: false,
		CreatedAt:  createdAt,
	}
	return user, us.userRepo.Create(user)
}

func (us *UserService) Update(id *uuid.UUID, params *dto.UpdateUserParams) (*model.User, error) {
	user, err := us.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if params.Username != nil {
		user.Username = *params.Username
	}

	if params.Password != nil {
		hashed, _ := bcrypt.GenerateFromPassword(
			[]byte(*params.Password),
			bcrypt.DefaultCost)
		user.Password = string(hashed)
	}

	if params.Email != nil {
		user.Email = *params.Email
	}

	if params.IsDisabled != nil {
		user.IsDisabled = *params.IsDisabled
	}

	return &user, us.userRepo.Update(&user)
}
