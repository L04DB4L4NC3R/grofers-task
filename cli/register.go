package cli

import (
	"fmt"
	"os"
	"os/exec"
)

/*
RegisterCLI :  Handler for command line args and work delegator
*/
func RegisterCLI() {

	// handler for args
	if len(os.Args) == 1 {
		// clear terminal
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
		fmt.Print("================================= grof CLI ==============================\n\n\n")
		fmt.Println("Usage: ")
		fmt.Println("grof put [key] [value] --------------------> Set key-value")
		fmt.Println("grof get [key] ----------------------------> Get value from key")
		fmt.Println("grof delete [key] -------------------------> Delete from DB")
		fmt.Println("grof update [key] [new value] -------------> Update DB with new value")
		fmt.Println("grof watch --------------------------------> Subscribe to changes in DB")
		fmt.Print("grof watch --------------------------------> Watch for realtime DB changes\n\n\n")
		os.Exit(1)
	} else if os.Args[1] == "put" && len(os.Args) >= 4 {
		putHandler()
	} else if os.Args[1] == "get" && len(os.Args) >= 3 {
		getHandler()
	} else if os.Args[1] == "delete" && len(os.Args) >= 3 {
		delHandler()
	} else if os.Args[1] == "update" && len(os.Args) >= 4 {
		updateHandler()
	} else if os.Args[1] == "watch" {
		watchHandler()
	} else {
		fmt.Println("\n\nUsage: ")
		fmt.Println("grof put [key] [value] --------------------> Set key-value")
		fmt.Println("grof get [key] ----------------------------> Get value from key")
		fmt.Println("grof delete [key] -------------------------> Delete from DB")
		fmt.Println("grof update [key] [new value] -------------> Update DB with new value")
		fmt.Println("grof watch --------------------------------> Subscribe to changes in DB")
		fmt.Print("grof watch --------------------------------> Watch for realtime DB changes\n\n\n")
		os.Exit(1)
	}
}
