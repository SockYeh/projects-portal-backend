package models

import "github.com/google/uuid"

type Invite struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Role      string    `gorm:"not null"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt int64     `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
