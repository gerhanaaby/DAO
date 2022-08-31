package models

import (
	"gorm.io/gorm"
)

type ApplicantDataBank struct {
	gorm.Model
	ID                     uint   `validate:"isdefault"`
	MonthlyDebitFreq       string `gorm:"size:256;not null;" validate:"required"`
	MonthlyCredFreq        string `gorm:"size:256;not null;" validate:"required"`
	MonthlyDebAmt          string `gorm:"size:256;not null;" validate:"required"`
	MonthlyCredAmt         string `gorm:"size:256;not null;" validate:"required"`
	Purpose                string `gorm:"size:256;not null;" validate:"required"`
	Source                 string `gorm:"size:256;not null;" validate:"required"`
	Checked                string `gorm:"size:256;not null;" validate:"required"`
	Version                string `gorm:"size:256;not null;" validate:"required"`
	VersionCustomerConsent string `gorm:"size:256;not null;" validate:"required"`
	CIFCode                string `gorm:"size:256;not null;" validate:"required"`
	Format                 string `gorm:"size:256;not null;" validate:"required"`
	OnSubmit               string `gorm:"size:256;not null;" validate:"required"`
	CustomerConsent        string `gorm:"size:256;not null;" validate:"required"`
}

//version
