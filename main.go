package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	"github.com/angadsharma1016/grofers-task/cli"
	"github.com/angadsharma1016/grofers-task/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

type server struct {
	DBCon *sql.DB
}

func main() {

	// setup a server instance
	var s server

	// load env config
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// connect to MySQL
	s.DBCon = model.ConnectDB()
	defer s.DBCon.Close()

	// register the CLI
	cli.RegisterCLI()

}
