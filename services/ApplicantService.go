package services

import (
	"DAO/database"
	"DAO/models"
	"fmt"
	"time"
)

func CreateApplicant(applicant models.ApplicantsRequest, KTP string, SelfieKTP string) error {
	var dataApplicant models.Applicants
	dataApplicant.Education = applicant.Education
	dataApplicant.Email = applicant.Email
	dataApplicant.Name = applicant.Name
	dataApplicant.HaveNPWP = applicant.HaveNPWP
	dataApplicant.Income = applicant.Income
	dataApplicant.Religion = applicant.Religion
	dataApplicant.JenisKelamin = applicant.JenisKelamin
	dataApplicant.KTPImage = KTP
	dataApplicant.SelfieImage = SelfieKTP
	dataApplicant.MaritalStatus = applicant.MaritalStatus
	dataApplicant.MotherName = applicant.MotherName
	dataApplicant.MobileNumber = applicant.MobileNumber
	dataApplicant.NPWP = applicant.NPWP
	dataApplicant.TmpLahir = applicant.TmpLahir
	dataApplicant.KTPNumber = applicant.KTPNumber
	date, error := time.Parse("2006-01-02", applicant.DOB)
	if error != nil {
		fmt.Println(error)
		return error
	}
	dataApplicant.DOB = date
	dataApplicant.ApplicantAlamat = applicant.ApplicantAlamat
	dataApplicant.ApplicantHistory.Notes = "-"
	dataApplicant.ApplicantHistory.Status = "Complete"
	dataApplicant.ApplicantHistory.Step = 1
	fmt.Println(dataApplicant)
	err := database.GetDB().Debug().Create(&dataApplicant).Error

	return err
}
