package cli

import (
	"fmt"
	"os"

	nats "github.com/nats-io/go-nats"
)

var conn *nats.Conn

func RegisterCLI(cn *nats.Conn) {
	conn = cn
	if len(os.Args) == 1 {
		fmt.Println("Usage: ")
		fmt.Println("./bin/main put [key] [value] --------------------> Set key-value")
		fmt.Println("./bin/main get [key] ----------------------------> Get value from key")
		fmt.Println("./bin/main watch --------------------------------> Watch for realtime DB changes")
	} else if os.Args[1] == "put" {
		putHandler()
	} else if os.Args[1] == "get" {
		getHandler()
	} else if os.Args[1] == "watch" {
		watchHandler()
	}
}
