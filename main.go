package main

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"

	"github.com/angadsharma1016/grofers-task/cli"
	"github.com/angadsharma1016/grofers-task/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	nats "github.com/nats-io/go-nats"
)

type server struct {
	NatsCon *nats.EncodedConn
	DBCon   *sql.DB
}

func (s *server) SetupNats() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalln(err)
	}
	s.NatsCon = nec
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

	// connect to NATS
	s.SetupNats()
	defer s.NatsCon.Close()

	// register the CLI
	cli.RegisterCLI(s.NatsCon)

}
