package main

import (
	"log"

	"github.com/angadsharma1016/grofers-task/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	nats "github.com/nats-io/go-nats"
)

type server struct {
	NatsCon *nats.EncodedConn
}

func (s *server) SetupNats() {
	nc, _ := nats.Connect("nats:4222")
	nec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalln(err)
	}
	s.NatsCon = nec
}

func main() {
	db := model.ConnectDB()
	defer db.Close()
	var s server
	s.SetupNats()
	defer s.NatsCon.Close()
	select {}
}
