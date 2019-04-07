package cli

import (
	"fmt"
	"os"

	"github.com/angadsharma1016/grofers-task/model"
	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
)

func putHandler() {

	s := &model.Store{
		Key:   os.Args[2],
		Value: os.Args[3],
	}
	c := make(chan error)
	go s.PutValue(c)

	if err := <-c; err != nil {
		fmt.Println("Error putting values: ", err.Error())
		os.Exit(1)

	}
	fmt.Printf("Created key = %s and value = %s\n", s.Key, s.Value)

	// publish event
	store := pb.Store{
		Key:   s.Key,
		Value: s.Value,
	}
	data, err := proto.Marshal(&store)
	if err != nil {
		fmt.Println("Error unmarshalling to protobuf: ", err.Error())
		os.Exit(1)
	}
	if err = conn.Publish("PUT", data); err != nil {
		fmt.Println("Error publishing event: ", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
