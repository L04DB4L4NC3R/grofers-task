package model

import (
	"fmt"

	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func (s *Store) PutValue(c chan error) {
	res, err := con.Exec("INSERT INTO STORE VALUES(?, ?)", s.Key, s.Value)
	if err != nil {
		c <- err
		return
	}
	fmt.Println(res.RowsAffected())
	c <- nil
	return
}

func (s *Store) GetValue(c chan error) {
	row := con.QueryRow("SELECT v FROM STORE WHERE k = ?", s.Key)
	if err := row.Scan(&s.Value); err != nil {
		c <- err
		return
	}
	c <- nil
	return
}

func GetAllValues(c chan StoreReturn) {

	var storeArr []Store
	rows, err := con.Query("SELECT k, v FROM STORE WHERE")
	if err != nil {
		c <- StoreReturn{nil, err}
		return
	}
	for rows.Next() {
		var s Store
		err = rows.Scan(&s.Key, &s.Value)
		if err != nil {
			c <- StoreReturn{nil, err}
			return
		}
		storeArr = append(storeArr, s)
	}

	c <- StoreReturn{storeArr, nil}
	return
}

func (s Store) Publish(subject string, con *nats.Conn) {
	store := pb.Store{
		Key:   s.Key,
		Value: s.Value,
	}
	data, err := proto.Marshal(&store)
	if err != nil {
		fmt.Println("Error publishing event", err.Error())
	}
	con.Publish(subject, data)
	fmt.Println("Published event")
}
