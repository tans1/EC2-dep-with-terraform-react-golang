package domain

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model

	ID          uint64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserID    uint64
	CreatedBy User `gorm:"foreignkey:UserID;references:ID;"`

	Comments []Comment `gorm:"foreignkey:BlogID;references:ID;"`
}
