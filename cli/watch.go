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

	// subscribe to PUT
	_, err := conn.Subscribe("PUT", func(msg *nats.Msg) {
		fmt.Println("Got a hit on PUT event")
		data := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &data); err != nil {
			fmt.Println("An error occurred while unmarshalling protobuf: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("An Entry was added into the DB at: ", time.Now().String())
		fmt.Printf("key: %s\nvalue:%s", data.Key, data.Value)
	})

	if err != nil {
		fmt.Println("Error subscribing to PUT event: ", err.Error())
		os.Exit(1)
	}

	// subscribe forever
	select {}
}
