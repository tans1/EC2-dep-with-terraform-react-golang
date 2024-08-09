package domain

import (
	"time"

	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model

	ID          uint64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserID    uint64
	CreatedBy User `gorm:"foreignkey:UserID;references:ID;"`

	CommentID uint64
	Comment   Comment `gorm:"foreignkey:CommentID;references:ID;"`

	BlogID uint64
	Blog   Blog `gorm:"foreignkey:BlogID;references:ID;"`
}
