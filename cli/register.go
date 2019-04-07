package cli

import (
	"fmt"
	"os"

	nats "github.com/nats-io/go-nats"
)

var conn *nats.EncodedConn

func RegisterCLI(cn *nats.EncodedConn) {
	conn = cn
	if len(os.Args) == 1 {
		fmt.Println("Usage: ")
		fmt.Println("./bin/main put [key] [value] --------------------> Set key-value")
		fmt.Println("./bin/main get [key] ----------------------------> Get value from key")
		fmt.Println("./bin/main watch --------------------------------> Watch for realtime DB changes")
		os.Exit(1)
	} else if os.Args[1] == "put" && len(os.Args) >= 4 {
		putHandler()
	} else if os.Args[1] == "get" && len(os.Args) >= 3 {
		getHandler()
	} else if os.Args[1] == "watch" {
		watchHandler()
	} else {
		fmt.Println("Usage: ")
		fmt.Println("./bin/main put [key] [value] --------------------> Set key-value")
		fmt.Println("./bin/main get [key] ----------------------------> Get value from key")
		fmt.Println("./bin/main watch --------------------------------> Watch for realtime DB changes")
		os.Exit(1)
	}
}
