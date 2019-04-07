package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/angadsharma1016/grofers-task/cli"
	"github.com/angadsharma1016/grofers-task/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	nats "github.com/nats-io/go-nats"
)

type server struct {
	NatsCon *nats.EncodedConn
}

func (s *server) SetupNats() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalln(err)
	}
	s.NatsCon = nec
	log.Println("Connected to NATS")
}

func main() {

	// load env config
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// connect to MySQL
	db := model.ConnectDB()
	defer db.Close()

	// setup a server instance
	var s server

	// connect to NATS
	s.SetupNats()
	defer s.NatsCon.Close()

	// register the CLI
	cli.RegisterCLI(s.NatsCon)
}
