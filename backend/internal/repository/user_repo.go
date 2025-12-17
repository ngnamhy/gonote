package repository

import (
	"gonote/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(id uuid.UUID) (model.User, error) {
	var user model.User
	err := r.db.
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
