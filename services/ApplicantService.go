package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateApplicant(applicant models.Applicants) error {
	err := database.GetDB().Debug().Create(&applicant).Error
	return err
}
