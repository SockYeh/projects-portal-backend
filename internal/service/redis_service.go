package service

import (
	"os"
	"time"

	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/models"
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RedisService struct {
	Repo *repository.RedisRepo
}

func NewRedisService(r *repository.RedisRepo) *RedisService {
	return &RedisService{Repo: r}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func (svc *RedisService) CreateRefreshToken(userid uuid.UUID) (string, error) {
	var refreshToken = models.UserRefreshToken{UserID: userid.String()}

	claims := &models.JwtUserRefreshToken{
		UserID: userid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTok, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	refreshToken.RefreshToken = refreshTok

	if err := svc.Repo.SaveRefreshToken(&refreshToken); err != nil {
		return "", err
	}
	return refreshTok, nil
}

func (svc *RedisService) IsValidRefreshToken(refreshToken string) bool {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &models.JwtUserRefreshToken{}, jwtKeyFunc)

	claim, ok := parsedRefreshToken.Claims.(*models.JwtUserRefreshToken)
	if !ok || !parsedRefreshToken.Valid {
		return false
	}

	userRefreshToken, err := svc.Repo.GetRefreshToken(claim.UserID)
	if err != nil {
		return false
	}

	return userRefreshToken.RefreshToken == refreshToken
}
