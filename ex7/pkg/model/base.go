package model

import (
	"time"

	"gorm.io/gorm"
)

// Base is the base model for all models.
type Base struct {
	ID        int            `json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
