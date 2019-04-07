package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/angadsharma1016/grofers-task/model"
	"github.com/angadsharma1016/grofers-task/pb"
	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/go-nats"
)

/**
*@api {post} /put put values
*@apiName put values
*@apiGroup all
*@apiParam {string} key key of the value
*@apiParam {string} value value of the key
*@apiParamExample {json} request-example
*{
*    "key": "dsad",
*    "value": "ddd"
*}
*@apiParamExample {json} response-example
*{
*    "key": "dsad",
*    "value": "ddd"
*}
**/
func putHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := &model.Store{}

		// get POSTed data
		json.NewDecoder(r.Body).Decode(&store)
		c := make(chan error)

		// Run DB func
		go store.PutValue(c)

		// Handler errors
		if err := <-c; err != nil {
			json.NewEncoder(w).Encode(struct {
				Err string
			}{err.Error()})
			return
		}

		// connect to NATS
		natsConn, err := nats.Connect("nats:4222")
		if err != nil {
			fmt.Println("Error connecting to messaging service")
			os.Exit(1)
		}

		// create a protobuf instance and marshal it
		s := pb.Store{
			Key:   store.Key,
			Value: store.Value,
		}
		data, err := proto.Marshal(&s)
		if err != nil {
			fmt.Println("Error unmarshalling to protobuf: ", err.Error())
			os.Exit(1)
		}

		// publish protobuf in the PUT event
		if err = natsConn.Publish("PUT", data); err != nil {
			fmt.Println("Error publishing event: ", err.Error())
			os.Exit(1)
		}

		natsConn.Close()

		// Output values
		json.NewEncoder(w).Encode(store)
		return
	}
}
