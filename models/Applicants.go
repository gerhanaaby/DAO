package models

import (
	"time"

	"gorm.io/gorm"
)

type Applicants struct {
	gorm.Model
	ID               uint              `validate:"isdefault"`
	Name             string            `gorm:"size:256;not null;"`
	Email            string            `gorm:"size:256;not null;" validate:"required"`
	TmpLahir         string            `gorm:"size:256;not null;" validate:"required"`
	DOB              time.Time         `gorm:"not null;" validate:"required"`
	MotherName       string            `gorm:"size:256;not null;" validate:"required"`
	KTPNumber        string            `gorm:"size:256;not null;" validate:"required; numeric"`
	JenisKelamin     string            `gorm:"size:256;not null;" validate:"required"`
	Religion         string            `gorm:"size:20;not null;" validate:"required"`
	Education        string            `gorm:"size:50;not null;" validate:"required"`
	Income           uint32            `gorm:"size:256;not null;" validate:"required"`
	MaritalStatus    string            `gorm:"size:256;not null;" validate:"required"`
	HaveNPWP         bool              `gorm:"size:1;not null;" validate:"required"`
	NPWP             string            `gorm:"size:50;not null;" validate:"required"`
	MobileNumber     string            `gorm:"size:256;not null;" validate:"required; numeric"`
	KTPImage         string            `gorm:"size:256;not null;" validate:"required"`
	SelfieImage      string            `gorm:"size:256;not null;" validate:"required"`
	ApplicantAlamat  []ApplicantAlamat `gorm:"foreignkey:ApplicantID"`
	ApplicantHistory ApplicantHistory  `gorm:"foreignkey:ApplicantID"`
}
