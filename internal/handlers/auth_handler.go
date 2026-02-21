package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/models"
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler(service *service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func (h *AuthHandler) createRefreshToken(userid uuid.UUID) (string, error) {
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

	if err := h.Service.RedisRepo.SaveRefreshToken(&refreshToken); err != nil {
		return "", err
	}
	return refreshTok, nil
}

func (h *AuthHandler) isValidRefreshToken(refreshToken string) bool {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &models.JwtUserRefreshToken{}, jwtKeyFunc)

	claim, ok := parsedRefreshToken.Claims.(*models.JwtUserRefreshToken)
	if !ok || !parsedRefreshToken.Valid {
		return false
	}

	userRefreshToken, err := h.Service.RedisRepo.GetRefreshToken(claim.UserID)
	if err != nil {
		return false
	}

	return userRefreshToken.RefreshToken == refreshToken
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *echo.Context) error {
	var body LoginBody
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "bad body parameters"})
	}

	user, err := h.Service.Login(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	refreshToken, err := h.createRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	accessToken, err := h.Service.CreateAccessToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"access_token": accessToken, "refresh_token": refreshToken})
}
