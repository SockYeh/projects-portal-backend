package models

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	MilestoneID uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"not null"`
	Description string
	AssignedTo  *uuid.UUID
	Priority    string
	Status      string

	CreatedAt int64 `gorm:"autoCreateTime"`
	UpdatedAt int64 `gorm:"autoUpdateTime"`
}
