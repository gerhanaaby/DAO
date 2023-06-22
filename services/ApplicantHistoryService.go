package services

import (
	"DAO/Database"
	models "DAO/Models/tables"
)

func CreateApplicantHistory(applicantHistory models.ApplicantHistory) error {
	err := Database.GetDB().Debug().Create(&applicantHistory).Error
	return err
}

func UpdateApplicantHistory(applicantID string, applicantHistory models.ApplicantHistory) error {
	err := Database.GetDB().Debug().Model(&models.Applicants{}).Where("ApplicantID = ?", applicantID).Updates(&applicantHistory).Error
	return err
}
