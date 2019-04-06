package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var con *sql.DB

func ConnectDB() *sql.DB {
	str := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", str)
	if err != nil {
		log.Println(err)
		log.Println("Error connecting to DB")
		os.Exit(1)
		return nil
	}
	log.Println("Connected to DB")

	con = db
	return con
}

type Store struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type StoreReturn struct {
	Stores []Store `json:"stores"`
	Err    error   `json:"err"`
}
