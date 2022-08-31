package database

import (
	"DAO/models"
	"DAO/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	dbServer *gorm.DB
	err      error
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

}
func ConnectDBServer() {
	host := utils.CallConfig.HostdbServer
	port := utils.CallConfig.PortdbServer
	user := utils.CallConfig.UserdbServer
	password := utils.CallConfig.PassworddbServer
	dbname := utils.CallConfig.DbnameServer

	sqlInfo := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30", user, password, host, port, dbname)
	fmt.Println(sqlInfo)
	dbServer, err = gorm.Open(sqlserver.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
func GetDB() *gorm.DB {
	return db
}

func GetDBServer() *gorm.DB {
	return dbServer
}

func DBMigrate() {
	for _, model := range models.RegisterModels() {
		err := db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database migrated successfully")

}
func DBDropTable() {
	for _, model := range models.RegisterModels() {
		err := db.Debug().Migrator().DropTable(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database drop all table successfully")

}
