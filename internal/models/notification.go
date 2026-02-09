package models

import "github.com/google/uuid"

type Notification struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	Message string
	Read    bool `gorm:"default:false"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
