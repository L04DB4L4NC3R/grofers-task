package cli

import (
	"fmt"
	"os"

	"github.com/angadsharma1016/grofers-task/model"
	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func updateHandler() {

	// create key value struct and update DB
	s := &model.Store{
		Key:   os.Args[2],
		Value: "",
	}

	c := make(chan error)
	go s.UpdateValue(os.Args[3], c)

	if err := <-c; err != nil {
		fmt.Println("Error updating values: ", err.Error())
		os.Exit(1)

	}
	fmt.Printf("Updated key = %s and value = %s\n", s.Key, s.Value)

	// connect to NATS
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to messaging service")
		os.Exit(1)
	}

	// create a protobuf instance and marshal it
	store := pb.Store{
		Key:   s.Key,
		Value: s.Value,
	}
	data, err := proto.Marshal(&store)
	if err != nil {
		fmt.Println("Error unmarshalling to protobuf: ", err.Error())
		os.Exit(1)
	}

	// publish protobuf in the UPDATE event
	if err = natsConn.Publish("UPDATE", data); err != nil {
		fmt.Println("Error publishing event: ", err.Error())
		os.Exit(1)
	}

	natsConn.Close()
	os.Exit(0)

}
