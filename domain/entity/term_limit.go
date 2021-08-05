package entity

import (
	"time"

	"gorm.io/gorm"
)

type TermLimit struct {
	ID               uint64 `gorm:"primary_key"` // primary key
	Amount             string
	CreatedAt        time.Time      // timestamp
	DeletedAt        gorm.DeletedAt // nullable timestamp (soft delete)
	UpdatedAt        time.Time      // timestamp
}