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
	PayerID     int //User ID
	RecieverID  int //User ID
	Description string
}

type TransactionPageRequest struct {
	ID    int    `query:"id"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Sort  string `query:"sort"`
}

type TransactionResponse struct {
	Amount      float64 `json:"amount"`
	PayerID     int     `json:"payer_id"`
	RecieverID  int     `json:"reciever_id"`
	Description string  `json:"description"`
}
