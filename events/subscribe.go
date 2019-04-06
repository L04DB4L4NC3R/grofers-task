package events

import (
	"fmt"

	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func Subscribe(subject string, con *nats.Conn) {
	con.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println("Subscribed to", subject)
		message := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &message); err != nil {
			fmt.Println("Error unmarshaling event response", err.Error())
			return
		}
		fmt.Printf("Key: %s 	Value: %s", message.Key, message.Value)
	})

	select {}
}
