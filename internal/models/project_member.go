package models

import "github.com/google/uuid"

type ProjectMember struct {
	ProjectID uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;primaryKey"`
	Role      string    // maintainer / member / viewer
}
