package model

import (
	"time"

	"gorm.io/gorm"
)

// Base is the base model for all models.
type Base struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
