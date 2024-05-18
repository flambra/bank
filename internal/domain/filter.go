package domain

import (
	"gorm.io/gorm"
)

type TransactionPageFilter struct {
	ID int `query:"id"`
}

func (f TransactionPageFilter) Apply(db *gorm.DB) *gorm.DB {
	return db.Where("payer_id = ? OR reciever_id = ?", f.ID, f.ID)
}
