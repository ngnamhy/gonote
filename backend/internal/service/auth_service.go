package service

import (
	"gonote/internal/dto"
	"gonote/internal/model"
	"gonote/internal/repository"
	"gonote/pkg/auth"
	"gonote/pkg/cache"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtService *auth.JWTService
	redisCache *cache.RedisCache
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService, redisCache *cache.RedisCache) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
		redisCache: redisCache}
}

func (as *AuthService) Login(params *dto.LoginParams) (*string, *string, *model.User, error) {
	user, err := as.userRepo.GetByEmail(&params.Email)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, nil, nil, err
	}

	accessToken, err := as.jwtService.GenAccessToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	refreshToken, err := as.jwtService.GenRefreshToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	err = as.jwtService.StoreRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, nil, err
	}

	return &accessToken, &refreshToken.Token, user, nil
}

func (as *AuthService) Register(params *dto.RegisterParams) (*string, *string, *model.User, error) {
	id := uuid.New()
	createdAt := time.Now()

	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(*params.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	user := &model.User{
		ID:         id,
		Username:   *params.Username,
		Password:   string(hashed),
		Email:      *params.Email,
		IsDisabled: false,
		CreatedAt:  createdAt,
		AvatarURL:  "",
	}
	err = as.userRepo.Create(user)
	if err != nil {
		return nil, nil, nil, err
	}

	accessToken, err := as.jwtService.GenAccessToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	refreshToken, err := as.jwtService.GenRefreshToken(user)
	if err != nil {
		return nil, nil, nil, err
	}

	err = as.jwtService.StoreRefreshToken(refreshToken)
	if err != nil {
		return nil, nil, nil, err
	}

	return &accessToken, &refreshToken.Token, user, nil
}

func (h *AuthService) Logout(c *gin.Context) {

}
