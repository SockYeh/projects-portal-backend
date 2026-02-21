package repository

import (
	"github.com/ComputerSocietyVITC/projects-portal-backend/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (r *AuthRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepo) GetUserByEmail(email string) (*models.User, error) {
	var user = models.User{Email: email}
	if err := r.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) GetUserByID(userUUID uuid.UUID) (*models.User, error) {
	var user = models.User{ID: userUUID}
	if err := r.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepo) GetUserRole(userUUID uuid.UUID) (string, error) {
	var roleID = models.UserRole{UserID: userUUID}
	if err := r.DB.First(&roleID).Error; err != nil {
		return "", err
	}

	var role = models.Role{ID: roleID.RoleID}
	if err := r.DB.First(&role).Error; err != nil {
		return "", err
	}

	return role.Name, nil
}

func (r *AuthRepo) CreateRoleReference(userUUID uuid.UUID, roleName string) error {
	var role = models.Role{Name: roleName}
	if err := r.DB.First(&role).Error; err != nil {
		return err
	}

	var userRole = models.UserRole{UserID: userUUID, RoleID: role.ID}
	return r.DB.Create(userRole).Error
}
