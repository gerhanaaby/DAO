package main

import (
	"DAO/database"
	"DAO/utils"
	"fmt"
)

func init() {
	utils.LoadConfig()
	database.ConnectDB()
}

func main() {

	fmt.Printf("helloworld")
}
