package models

import "github.com/google/uuid"

type Milestone struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	DueDate   string
	Progress  int `gorm:"default:0"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
