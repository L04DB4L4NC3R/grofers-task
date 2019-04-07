package cli

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/angadsharma1016/grofers-task/model"
)

func getHandler() {
	s := &model.Store{
		Key:   os.Args[2],
		Value: "",
	}
	c := make(chan error)
	go s.GetValue(c)

	if err := <-c; err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No values found for key: ", s.Key)
			return
		}

		fmt.Println(err.Error())
		return

	}
	fmt.Println(s.Value)
	return
}
