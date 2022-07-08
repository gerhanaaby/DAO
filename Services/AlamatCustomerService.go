package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateApplicantAlamat(applicantAlamat models.ApplicantAlamat) error {
	err := database.GetDB().Debug().Create(&applicantAlamat).Error
	return err
}
