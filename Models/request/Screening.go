package models

import (
	"gorm.io/gorm"
)

type Screening struct {
	gorm.Model
	ID             uint   `validate:"isdefault"`
	CIFNumber      string `gorm:"size:256;not null;" validate:"required" binding:"required"`
	Name           string `gorm:"size:256;not null;" validate:"required" binding:"required"`
	Nationality    string `gorm:"size:256;not null;" validate:"required" binding:"required"`
	DOB            string `gorm:"not null;" validate:"required" binding:"required"`
	MatchScore     string `gorm:"size:256;not null;" validate:"required" binding:"required"`
	IdentityNumber string `gorm:"size:256;not null;" validate:"required" binding:"required"`
}
