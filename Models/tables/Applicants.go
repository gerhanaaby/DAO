package models

import (
	"time"

	"gorm.io/gorm"
)

type Applicants struct {
	gorm.Model
	ID               uint              `validate:"isdefault"`
	CIFNumber        string            `gorm:"size:256;not null;" validate:"required"`
	Name             string            `gorm:"size:256;not null;" validate:"required"`
	Email            string            `gorm:"size:256;not null;" validate:"required"`
	TmpLahir         string            `gorm:"size:256;not null;" validate:"required"`
	DOB              time.Time         `gorm:"not null;" validate:"required"`
	MotherName       string            `gorm:"size:256;not null;" validate:"required"`
	KTPNumber        string            `gorm:"size:256;not null;" validate:"required; numeric"`
	JenisKelamin     string            `gorm:"size:256;not null;" validate:"required"`
	Religion         string            `gorm:"size:20;not null;" validate:"required"`
	Education        string            `gorm:"size:50;not null;" validate:"required"`
	Income           uint64            `gorm:"size:256;not null;" validate:"required"`
	MaritalStatus    string            `gorm:"size:256;not null;" validate:"required"`
	HaveNPWP         bool              `gorm:"size:1;not null;" validate:"required"`
	MobileNumber     string            `gorm:"size:256;not null;" validate:"required; numeric"`
	KTPImage         string            `gorm:"size:256;not null;" validate:"required"`
	SelfieImage      string            `gorm:"size:256;not null;" validate:"required"`
	SelfieScore      int               `gorm:"not null;" validate:"required"`
	ApplicantAlamat  []ApplicantAlamat `gorm:"foreignkey:ApplicantID"`
	ApplicantOffice  ApplicantOffice   `gorm:"foreignkey:ApplicantID"`
	ApplicantHistory ApplicantHistory  `gorm:"foreignkey:ApplicantID"`
	Pekerjaan        string            `gorm:"size:256;not null;" validate:"required"`
}

type ApplicantRequests struct {
	gorm.Model
	ID               uint               `validate:"isdefault"`
	Name             string             `gorm:"size:256;not null;" validate:"required"`
	Email            string             `gorm:"size:256;not null;" validate:"required"`
	TmpLahir         string             `gorm:"size:256;not null;" validate:"required"`
	DOB              time.Time          `gorm:"not null;" validate:"required"`
	MotherName       string             `gorm:"size:256;not null;" validate:"required"`
	KTPNumber        string             `gorm:"size:256;not null;" validate:"required; numeric"`
	JenisKelamin     string             `gorm:"size:256;not null;" validate:"required"`
	Religion         string             `gorm:"size:20;not null;" validate:"required"`
	Education        string             `gorm:"size:50;not null;" validate:"required"`
	Income           uint64             `gorm:"size:256;not null;" validate:"required"`
	MaritalStatus    string             `gorm:"size:256;not null;" validate:"required"`
	HaveNPWP         bool               `gorm:"size:1;not null;" validate:"required"`
	MobileNumber     string             `gorm:"size:256;not null;" validate:"required; numeric"`
	KTPImage         string             `gorm:"size:256;not null;" validate:"required"`
	SelfieImage      string             `gorm:"size:256;not null;" validate:"required"`
	SelfieScore      int                `gorm:"not null;" validate:"required"`
	ApplicantAlamat  [1]ApplicantAlamat `gorm:"foreignkey:ApplicantID"`
	ApplicantOffice  ApplicantOffice    `gorm:"foreignkey:ApplicantID"`
	ApplicantHistory ApplicantHistory   `gorm:"foreignkey:ApplicantID"`
	Pekerjaan        string             `gorm:"size:256;not null;" validate:"required"`
}

var AppQ []Applicants

// type ApplicantsRequest struct {
// 	gorm.Model
// 	ID               uint                  `validate:"isdefault"`
// 	Name             string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	Email            string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	TmpLahir         string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	DOB              string                `gorm:"not null;" validate:"required" binding:"required"`
// 	MotherName       string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	KTPNumber        string                `gorm:"size:256;not null;" validate:"required; numeric" binding:"required"`
// 	JenisKelamin     string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	Religion         string                `gorm:"size:20;not null;" validate:"required" binding:"required"`
// 	Education        string                `gorm:"size:50;not null;" validate:"required" binding:"required"`
// 	Income           uint32                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	MaritalStatus    string                `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	HaveNPWP         bool                  `gorm:"size:1;not null;" validate:"required" binding:"required"`
// 	NPWP             string                `gorm:"size:50;not null;" validate:"required" binding:"required"`
// 	MobileNumber     string                `gorm:"size:256;not null;" validate:"required; numeric" binding:"required"`
// 	KTPImage         *multipart.FileHeader `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	SelfieImage      *multipart.FileHeader `gorm:"size:256;not null;" validate:"required" binding:"required"`
// 	ApplicantAlamat  []ApplicantAlamat     `gorm:"foreignkey:ApplicantID"`
// 	ApplicantHistory ApplicantHistory      `gorm:"foreignkey:ApplicantID"`
// }
