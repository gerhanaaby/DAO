package Models

import (
	"gorm.io/gorm"
)

type AlamatCustomer struct {
	gorm.Model
	ID         uint   `validate:"isdefault"`
	Alamat     string `gorm:"size:256;not null;unique" validate:"isdefault"`
	Kecamatan  string `gorm:"size:256;not null;" validate:"required"`
	Provinsi   string `gorm:"size:256;not null;" validate:"required"`
	Kota       string `gorm:"size:256;not null;" validate:"required"`
	RT         string `gorm:"size:256;not null;" validate:"required"`
	RW         string `gorm:"size:256;not null;" validate:"required"`
	KodePos    string `gorm:"size:256;not null;" validate:"required"`
	CustomerID string `gorm:"size:256;not null;" validate:"required"`
}
