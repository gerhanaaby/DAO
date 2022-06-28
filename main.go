package main

import (
	database "DAO/Database"
	utils "DAO/Utils"
	"fmt"
)

func init() {
	utils.LoadConfig()
	database.ConnectDB()
}

func main() {

	fmt.Printf("helloworld")
}
