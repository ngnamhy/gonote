package auth

import (
	"crypto/rand"
	"encoding/base64"
	"gonote/internal/model"
	"gonote/pkg/cache"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	redisCache *cache.RedisCache
}

func NewJWTService(redisCache *cache.RedisCache) *JWTService {
	return &JWTService{
		redisCache: redisCache,
	}
}

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

const (
	AccessTokenTTL  = 15 * time.Minute
	RefreshTokenTTL = 1 * 24 * time.Hour
)

func (jwtService *JWTService) GenAccessToken(user *model.User) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "backend-gonote",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Print(string(jwtSecret))
	return token.SignedString(jwtSecret)
}

type RefreshToken struct {
	Token     string    `json:"token"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	Revoked   bool      `json:"revoked"`
}

func (jwtService *JWTService) GenRefreshToken(user *model.User) (RefreshToken, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return RefreshToken{}, err
	}
	Token := base64.URLEncoding.EncodeToString(bytes)

	return RefreshToken{
		Token:     Token,
		UserID:    user.ID.String(),
		ExpiresAt: time.Now().Add(RefreshTokenTTL),
		Revoked:   false,
	}, nil
}

func (jwtService *JWTService) StoreRefreshToken(token RefreshToken) error {
	err := jwtService.redisCache.Set(token.Token, token, AccessTokenTTL)
	return err
}
