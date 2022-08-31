package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Port             string
	Hostdb           string
	Portdb           string
	Userdb           string
	Passworddb       string
	Dbname           string
	PortServer       string
	HostdbServer     string
	PortdbServer     string
	UserdbServer     string
	PassworddbServer string
	DbnameServer     string
}

var CallConfig *DBConfig

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	CallConfig = &DBConfig{
		Port:             os.Getenv("PORT"),
		Hostdb:           os.Getenv("HOSTDB"),
		Portdb:           os.Getenv("PORTDB"),
		Userdb:           os.Getenv("USERDB"),
		Passworddb:       os.Getenv("PASSWORDDB"),
		Dbname:           os.Getenv("DBNAME"),
		PortServer:       os.Getenv("PORT_SERVER"),
		HostdbServer:     os.Getenv("HOSTDB_SERVER"),
		PortdbServer:     os.Getenv("PORTDB_SERVER"),
		UserdbServer:     os.Getenv("USERDB_SERVER"),
		PassworddbServer: os.Getenv("PASSWORDDB_SERVER"),
		DbnameServer:     os.Getenv("DBNAME_SERVER"),
	}
}
