package models

import (
	"gorm.io/gorm"
)

type ApplicantAlamat struct {
	gorm.Model
	ID          uint   `validate:"isdefault"`
	Alamat      string `gorm:"size:256;not null;unique" validate:"isdefault"`
	Kecamatan   string `gorm:"size:;not null;" validate:"required"`
	Provinsi    string `gorm:"size:100;not null;" validate:"required"`
	Kota        string `gorm:"size:100;not null;" validate:"required"`
	RT          uint16 `gorm:"size:5;not null;" validate:"required"`
	RW          uint16 `gorm:"size:5;not null;" validate:"required"`
	KodePos     uint16 `gorm:"size:5;not null;" validate:"required"`
	ApplicantID uint   `gorm:"size:256;not null;" validate:"required"`
}
