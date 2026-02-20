package handlers

import (
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/service"
	"go.uber.org/zap"
)

type Handlers struct {
	Auth   *AuthHandler
	Invite *InviteHandler
}

func NewHandlers(authService *service.AuthService, inviteService *service.InviteService, redisService *service.RedisService, logger *zap.Logger) *Handlers {
	return &Handlers{
		Auth: NewAuthHandler(authService, redisService, logger),
		// Invite: NewInviteHandler(inviteService, logger),
	}
}
