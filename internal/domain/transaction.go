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
	Amount      int
	PayerID     int //User ID
	RecieverID  int //User ID
	Description string
}

type TransactionCreateRequest struct {
	Amount      int    `json:"amount"`
	PayerID     int    `json:"payer_id"`
	RecieverID  int    `json:"reciever_id"`
	Description string `json:"description"`
}

type TransactionPageRequest struct {
	ID    int    `query:"id"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Sort  string `query:"sort"`
}

type TransactionPageResponse struct {
	Amount      int `json:"amount"`
	PayerID     int     `json:"payer_id"`
	RecieverID  int     `json:"reciever_id"`
	Description string  `json:"description"`
}

type TransactionPageFilter struct {
	ID int `query:"id"`
}

func (f TransactionPageFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where("payer_id = ? OR reciever_id = ?", f.ID, f.ID)
}
