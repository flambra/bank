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

type InstitutionCreateRequest struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
	Logo string `json:"logo"`
}

type InstitutionUpdateRequest struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
	Logo string `json:"logo"`
}

type InstitutionListResponse struct {
	ID   uint   `json:"id"`
	Code int    `json:"code"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
