package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Port       string
	Hostdb     string
	Portdb     string
	Userdb     string
	Passworddb string
	Dbname     string
}

var CallConfig *DBConfig

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	CallConfig = &DBConfig{
		Port:       os.Getenv("PORT"),
		Hostdb:     os.Getenv("HOSTDB"),
		Portdb:     os.Getenv("PORTDB"),
		Userdb:     os.Getenv("USERDB"),
		Passworddb: os.Getenv("PASSWORDDB"),
		Dbname:     os.Getenv("DBNAME"),
	}
}
