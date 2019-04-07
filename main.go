package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/angadsharma1016/grofers-task/cli"
	"github.com/angadsharma1016/grofers-task/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/rs/cors"             // for godep cors
)

func main() {

	// load env config
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// connect to MySQL
	db := model.ConnectDB()
	defer db.Close()

	// register the CLI
	cli.RegisterCLI()

}
