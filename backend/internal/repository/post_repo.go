package repository

import (
	"gonote/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAllPost() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *PostRepository) GetByID(id uuid.UUID) (model.Post, error) {
	var post model.Post
	err := r.db.
		Where("id = ?", id).
		First(&post).
		Error

	if err != nil {
		return model.Post{}, err
	}

	return post, nil

}

func (r *PostRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) Update(post *model.Post) error {
	err := r.db.Save(post).Error
	return err
}
