package service

import (
	"os"
	"strings"
	"time"

	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/models"
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.AuthRepo
}

func NewAuthService(r *repository.AuthRepo) *AuthService {
	return &AuthService{Repo: r}
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
	parsedRefreshToken, _ := jwt.ParseWithClaims(accessToken, &models.JwtUserAccessToken{}, jwtKeyFunc)

	if _, ok := parsedRefreshToken.Claims.(*models.JwtUserAccessToken); !ok && !parsedRefreshToken.Valid {
		return false
	}

	return true
}
