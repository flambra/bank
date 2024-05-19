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

type AccountCreateRequest struct {
	Account       int    `json:"account"`
	Balance       int    `json:"balance"`
	Agency        int    `json:"agency"`
	Digit         int    `json:"digit"`
	Favorite      bool   `json:"favorite"`
	Owner         string `json:"owner"`
	TypeAccount   string `json:"type_account"`
	InstitutionID int    `json:"institution_id"`
}

type AccountUpdateRequest struct {
	Account       int    `json:"account"`
	Balance       int    `json:"balance"`
	Agency        int    `json:"agency"`
	Digit         int    `json:"digit"`
	Favorite      bool   `json:"favorite"`
	Owner         string `json:"owner"`
	TypeAccount   string `json:"type_account"`
	InstitutionID int    `json:"institution_id"`
}
