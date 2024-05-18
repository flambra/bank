package domain

import (
	"time"

	"gorm.io/gorm"
)

type Institution struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Code      int            `gorm:"unique"`
	Name      string         `gorm:"not null"`
	CNPJ      string
	Logo      string
}
