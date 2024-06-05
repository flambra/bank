package domain

import (
	"time"

	"gorm.io/gorm"
)

type BalanceHistorical struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BalanceID uint           `gorm:"not null;index"`
	AccountID uint           `gorm:"not null;index"`
	Amount    int        `gorm:"not null"`
}

type BalanceCurrent struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	AccountID uint           `gorm:"not null;index"`
	Amount    int        `gorm:"not null"`
}


type BalanceUpdateRequest struct {
	Amount int `json:"amount"`
}