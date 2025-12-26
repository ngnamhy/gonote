package service

import (
	"gonote/internal/dto"
	"gonote/internal/repository"
	"gonote/pkg/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtService *auth.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService) *AuthService {
	return &AuthService{userRepo: userRepo, jwtService: jwtService}
}

func (as *AuthService) Login(params dto.LoginParams) (*string, error) {
	user, err := as.userRepo.GetByUsername(params.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
		return nil, err
	}

	accessToken, err := as.jwtService.GenAccessToken(user)

	if err != nil {
		return nil, err
	}
	return &accessToken, nil
}

func (h *AuthService) Logout(c *gin.Context) {

}
