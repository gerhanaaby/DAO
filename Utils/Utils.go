package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	Hostdb     string
	Portdb     string
	Userdb     string
	Passworddb string
	Dbname     string
}

var CallConfig *Config

func LoadConfig() {
	err := godotenv.Load("./resources/.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	CallConfig = &Config{
		Port:       os.Getenv("PORT"),
		Hostdb:     os.Getenv("HOSTDB"),
		Portdb:     os.Getenv("PORTDB"),
		Userdb:     os.Getenv("USERDB"),
		Passworddb: os.Getenv("PASSWORDDB"),
		Dbname:     os.Getenv("DBNAME"),
	}
}
