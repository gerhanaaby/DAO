package models

import (
	"gorm.io/gorm"
)

type ApplicantOffice struct {
	gorm.Model
	ID                uint   `validate:"isdefault"`
	OfficeName        string `gorm:"size:256;not null;" validate:"isdefault"`
	OfficeAddress     string `gorm:"size:256;not null;" validate:"required"`
	OfficeZipCode     string `gorm:"size:256;not null;" validate:"required"`
	OfficePhoneNumber string `gorm:"size:256;not null;" validate:"required"`
	Pekerjaan         string `gorm:"size:256;not null;" validate:"required"`
	ApplicantID       string `gorm:"size:256;not null;" validate:"required"`
}
