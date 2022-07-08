package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateApplicant(applicant models.ApplicantsRequest) error {
	err := database.GetDB().Debug().Table("applicants").Create(&applicant).Error
	return err
}
