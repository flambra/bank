package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Amount      float64
	Payer       int
	Reciever    int
	Description string
}
