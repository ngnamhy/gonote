package service

import (
	"gonote/internal/dto"
	"gonote/internal/model"
	"gonote/internal/repository"
	"time"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo *repository.PostRepository
}

func NewPostService(postRepo *repository.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (ps *PostService) GetAllPost() ([]model.Post, error) {
	posts, err := ps.postRepo.GetAllPost()
	return posts, err
}

func (ps *PostService) GetByID(id uuid.UUID) (model.Post, error) {
	return ps.postRepo.GetByID(id)
}

func (ps *PostService) Create(params dto.CreatePostParams) (*model.Post, error) {
	id := uuid.New()
	createdAt := time.Now()
	post := &model.Post{
		ID:        id,
		CreatedAt: createdAt,
		UserID:    params.UserID,
		Title:     params.Title,
		Body:      params.Body,
		Score:     0,
	}
	err := ps.postRepo.Create(post)
	return post, err
}

func (ps *PostService) Update(id uuid.UUID, params dto.UpdatePostParams) (*model.Post, error) {
	post, err := ps.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if params.Title != nil {
		post.Title = *params.Title
	}

	if params.Body != nil {
		post.Body = params.Body
	}

	if params.Score != nil {
		post.Score = *params.Score
	}

	return &post, ps.postRepo.Update(&post)
}
