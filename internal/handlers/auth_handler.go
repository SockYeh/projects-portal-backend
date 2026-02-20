package handlers

import (
	"net/http"

	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/service"
	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service  *service.AuthService
	RedisSvc *service.RedisService
	Logger   *zap.Logger
}

func NewAuthHandler(service *service.AuthService, redisSvc *service.RedisService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		Service:  service,
		RedisSvc: redisSvc,
		Logger:   logger,
	}
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
	h.Logger.Log(zap.InfoLevel, "Login attempt", zap.String("email", body.Email), zap.String("password", body.Password))

	user, err := h.Service.Login(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	refreshToken, err := h.RedisSvc.CreateRefreshToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	accessToken, err := h.Service.CreateAccessToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"access_token": accessToken, "refresh_token": refreshToken})
}
