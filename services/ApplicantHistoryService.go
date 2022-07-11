package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateApplicantHistory(applicantHistory models.ApplicantHistory) error {
	applicantHistory.Notes = "-"
	applicantHistory.Status = "Complete"
	applicantHistory.Step = 1
	applicantHistory.ApplicantID = 1
	err := database.GetDB().Debug().Create(&applicantHistory).Error
	return err
}

func UpdateApplicantHistory(applicantID string, applicantHistory models.ApplicantHistory) error {

	err := database.GetDB().Debug().Model(&models.Applicants{}).Where("ApplicantID = ?", applicantID).Updates(&applicantHistory).Error
	return err
}
