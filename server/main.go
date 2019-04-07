package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/angadsharma1016/grofers-task/model"
	"github.com/angadsharma1016/grofers-task/server/controller"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

type server struct {
	Host   string
	Port   string
	DBConn *sql.DB
}

func (s server) Run() {
	h := controller.Startup()
	log.Println("Listening....")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), *h))
}

func main() {

	s := server{"0.0.0.0", "3000", nil}

	s.DBConn = model.ConnectAPIToDB()
	defer s.DBConn.Close()
	s.Run()
}
