package models

import (
	"gorm.io/gorm"
)

type ApplicantHistory struct {
	gorm.Model
	ID          uint   `validate:"isdefault"`
	ApplicantID uint   `gorm:"size:256;not null;" validate:"required"`
	Step        int8   `gorm:"size:256;not null;" validate:"required"`
	Status      string `gorm:"size:256;not null;" validate:"required"`
	Notes       string `gorm:"size:256;not null;" validate:"required"`
}
