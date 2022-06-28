package routes

import (
	"DAO/database"
	"DAO/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func initCommands() {
	database.ConnectDB()
	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				database.DBMigrate()
				return nil
			},
		},
		{
			Name: "db:delete",
			Action: func(c *cli.Context) error {
				database.DBDropTable()
				return nil
			},
		},
		// {
		// 	Name: "db:seed",
		// 	Action: func(c *cli.Context) error {
		// 		err := seeders.DBSeed(server.DB)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		return nil
		// 	},
		// },
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func Runing(addr string) {
	fmt.Println("Welcome to DAO")
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
func Run() {

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		initCommands()
	} else {
		//Initialize()
		database.DBMigrate()
		Runing(":" + utils.CallConfig.Port)

	}
}
