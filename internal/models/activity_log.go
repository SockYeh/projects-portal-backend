package models

import (
	"github.com/google/uuid"
)

type ActivityLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    *uuid.UUID
	Action    string    `gorm:"not null"`
	Entity    string    `gorm:"not null"`
	EntityID  uuid.UUID `gorm:"not null"`
	CreatedAt int64     `gorm:"autoCreateTime"`
}
