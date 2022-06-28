package database

import (
	utils "DAO/Utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	host := utils.CallConfig.Hostdb
	port := utils.CallConfig.Portdb
	user := utils.CallConfig.Userdb
	password := utils.CallConfig.Passworddb
	dbname := utils.CallConfig.Dbname

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// db.Debug().AutoMigrate(&models.Customer{}, &models.Order{}, &models.Item{})
}
func GetDB() *gorm.DB {
	return db
}
