package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func watchHandler() {

	// connect to NATS
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to messaging service")
		os.Exit(1)
	}
	defer natsConn.Close()

	// subscribe to PUT
	natsConn.Subscribe("PUT", func(msg *nats.Msg) {
		fmt.Print("\n\n================= DATA ADDED IN DB =====================\n\n")
		// unmarshal protobuf
		data := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &data); err != nil {
			fmt.Println("An error occurred while unmarshalling protobuf: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("\nAn Entry was PUT into the DB at: ", time.Now().String())
		fmt.Printf("\nKey: %s\nValue:%s\n\n", data.Key, data.Value)
	})

	// subscribe to DELETE
	natsConn.Subscribe("DELETE", func(msg *nats.Msg) {
		fmt.Print("\n\n================= DATA REMOVED FROM DB =====================\n\n")
		// unmarshal protobuf
		data := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &data); err != nil {
			fmt.Println("An error occurred while unmarshalling protobuf: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("\nAn Entry was DELETED from the DB at: ", time.Now().String())
		fmt.Printf("\nKey: %s\nValue:%s\n\n", data.Key, data.Value)
	})

	// subscribe to UPDATE
	natsConn.Subscribe("UPDATE", func(msg *nats.Msg) {
		fmt.Print("\n\n================= DATA UPDATED IN DB =====================\n\n")
		// unmarshal protobuf
		data := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &data); err != nil {
			fmt.Println("An error occurred while unmarshalling protobuf: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("\nAn Entry was UPDATED from the DB at: ", time.Now().String())
		fmt.Printf("\nKey: %s\nNew value:%s\n\n", data.Key, data.Value)
	})

	// subscribe forever
	select {}
}
