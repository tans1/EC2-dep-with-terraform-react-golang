package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	ID          uint64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserID    uint64
	CreatedBy User `gorm:"foreignkey:UserID;references:ID;"`

	BlogID uint64
	Blog   Blog `gorm:"foreignkey:BlogID;references:ID;"`

	Reply []Reply `gorm:"foreignkey:CommentID;references:ID;"`
}
