package service

import (
	"gonote/internal/dto"
	"gonote/internal/model"
	"gonote/internal/repository"
	"gonote/pkg/auth"
	"gonote/pkg/cache"

	"github.com/gin-gonic/gin"
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

func (as *AuthService) Login(params dto.LoginParams) (*string, *string, *model.User, error) {
	user, err := as.userRepo.GetByUsername(params.Username)
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

func (h *AuthService) Logout(c *gin.Context) {

}
