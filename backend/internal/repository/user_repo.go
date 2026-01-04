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

func (r *UserRepository) GetAllUser() ([]model.User, error) {
	var users []model.User
	err := r.db.
		Find(&users).
		Where("is_disabled = fasle").
		Error

	return users, err
}

func (r *UserRepository) GetByID(id *uuid.UUID) (model.User, error) {
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

func (r *UserRepository) GetByUsername(username *string) (*model.User, error) {
	var user model.User
	err := r.db.
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email *string) (*model.User, error) {
	var user model.User
	err := r.db.
		Where("email = ?", email).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *model.User) error {
	err := r.db.Save(user).Error
	return err
}
