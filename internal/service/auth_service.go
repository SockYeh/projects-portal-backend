package service

import (
	"os"
	"strings"
	"time"

	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/models"
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo      *repository.AuthRepo
	RedisRepo *repository.RedisRepo
}

func NewAuthService(r *repository.AuthRepo, redis *repository.RedisRepo) *AuthService {
	return &AuthService{Repo: r, RedisRepo: redis}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func (svc *AuthService) CreateRefreshToken(userid uuid.UUID) (string, error) {
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

	if err := svc.RedisRepo.SaveRefreshToken(&refreshToken); err != nil {
		return "", err
	}
	return refreshTok, nil
}

func (svc *AuthService) IsValidRefreshToken(refreshToken string) bool {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &models.JwtUserRefreshToken{}, jwtKeyFunc)

	claim, ok := parsedRefreshToken.Claims.(*models.JwtUserRefreshToken)
	if !ok || !parsedRefreshToken.Valid {
		return false
	}

	userRefreshToken, err := svc.RedisRepo.GetRefreshToken(claim.UserID)
	if err != nil {
		return false
	}

	return userRefreshToken.RefreshToken == refreshToken
}

func (svc *AuthService) Login(email, password string) (*models.User, error) {
	user, err := svc.Repo.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(user.PasswordHash)), []byte(strings.TrimSpace(password)))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *AuthService) CreateAccessToken(user *models.User) (string, error) {
	var userid = user.ID

	role, err := svc.Repo.GetUserRole(userid)
	if err != nil {
		return "", err
	}

	claims := &models.JwtUserAccessToken{
		UserID: userid.String(),
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return accessToken, nil
}

func (svc *AuthService) IsValidAccessToken(accessToken string) bool {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &models.JwtUserAccessToken{}, jwtKeyFunc)

	if _, ok := parsedAccessToken.Claims.(*models.JwtUserAccessToken); !ok || !parsedAccessToken.Valid {
		return false
	}

	return true
}
