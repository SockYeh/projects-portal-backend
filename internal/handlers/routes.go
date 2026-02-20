package handlers

import (
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/repository"
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/service"
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, redis *redis.Client, logger *zap.Logger) {
	// inviteRepo := &repository.InviteRepo{DB: db, Logger: logger}
	authRepo := &repository.AuthRepo{DB: db, Logger: logger}
	redisRepo := &repository.RedisRepo{Redis: redis, Logger: logger}

	// inviteService := &service.InviteService{Repo: inviteRepo}
	authService := &service.AuthService{Repo: authRepo}
	redisService := &service.RedisService{Repo: redisRepo}

	handlers := NewHandlers(authService, nil, redisService, logger)

	e.POST("/auth/login", handlers.Auth.Login)
}
