package domain

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID            uint `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Account       int
	Balance       int
	Agency        int
	Digit         int
	Favorite      bool
	Owner         string
	TypeAccount   string
	InstitutionID int
}
