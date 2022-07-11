package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateApplicant(applicant models.Applicants) (models.Applicants, error) {
	applicant.ApplicantHistory.Notes = "-"
	applicant.ApplicantHistory.Status = "Complete"
	applicant.ApplicantHistory.Step = 1
	applicant.DOB = applicant.DOB.GoString()
	err := database.GetDB().Debug().Create(&applicant).Error
	return applicant, err
}
