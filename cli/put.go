package cli

import (
	"fmt"
	"os"

	"github.com/angadsharma1016/grofers-task/model"
)

func putHandler() {
	s := &model.Store{
		Key:   os.Args[2],
		Value: os.Args[3],
	}
	c := make(chan error)
	go s.PutValue(c)

	if err := <-c; err != nil {
		fmt.Println(err.Error())
		os.Exit(1)

	}
	fmt.Printf("Created key = %s and value = %s", s.Key, s.Value)
	os.Exit(0)
}
