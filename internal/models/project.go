package models

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"not null"`
	Description string
	Domain      string
	Status      string
	RepoURL     string
	StartDate   string
	EndDate     string
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null"`

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
