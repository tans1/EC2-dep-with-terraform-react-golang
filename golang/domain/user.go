package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID          uint64
	FirstName   string
	LastName    string
	Username    string
	Email       string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Blogs []Blog `gorm:"foreignkey:UserID;references:ID;"`
}
