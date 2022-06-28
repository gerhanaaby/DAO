package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID             uint             `validate:"isdefault"`
	Name           string           `gorm:"size:256;not null;unique" validate:"isdefault"`
	Email          string           `gorm:"size:256;not null;" validate:"required"`
	TmpLahir       string           `gorm:"size:256;not null;" validate:"required"`
	DOB            string           `gorm:"size:256;not null;" validate:"required"`
	MotherName     string           `gorm:"size:256;not null;" validate:"required"`
	KTPNumber      string           `gorm:"size:256;not null;" validate:"required"`
	JenisKelamin   string           `gorm:"size:256;not null;" validate:"required"`
	NPWP           string           `gorm:"size:256;not null;" validate:"required"`
	Religion       string           `gorm:"size:256;not null;" validate:"required"`
	Education      string           `gorm:"size:256;not null;" validate:"required"`
	Income         string           `gorm:"size:256;not null;" validate:"required"`
	MaritalStatus  string           `gorm:"size:256;not null;" validate:"required"`
	HaveNPWP       string           `gorm:"size:256;not null;" validate:"required"`
	MobileNumber   string           `gorm:"size:256;not null;" validate:"required"`
	KTPImage       string           `gorm:"size:256;not null;" validate:"required"`
	SelfieImage    string           `gorm:"size:256;not null;" validate:"required"`
	AlamatCustomer []AlamatCustomer `gorm:"foreignkey:CustomerID"`
	OfficeCustomer []OfficeCustomer `gorm:"foreignkey:CustomerID"`
}
