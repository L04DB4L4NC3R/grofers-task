package events

import (
	"fmt"
	"log"

	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

func Subscribe(subject string, con *nats.Conn) {
	con.Subscribe(subject, func(msg *nats.Msg) {
		message := pb.Store{}
		if err := proto.Unmarshal(msg.Data, &message); err != nil {
			log.Println("Error unmarshaling event response", err.Error())
			return
		}
		log.Println(fmt.Sprintf("Key: %s 	Value: %s", message.Key, message.Value))
	})

	select {}
}
