package services

import (
	"DAO/database"
	"DAO/models"
)

func CreateCustomer(customer models.Customer) error {
	err := database.GetDB().Debug().Create(&customer).Error
	return err
}
