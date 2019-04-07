package controller

import (
	"encoding/json"
	"net/http"

	"github.com/angadsharma1016/grofers-task/model"
)

/**
*@api {post} /get get values
*@apiName get values
*@apiGroup all
*@apiParam {string} key key of the value
*@apiParamExample {json} request-example
*{
*    "key": "dsad"
*}
*@apiParamExample {json} response-example
*{
*    "key": "dsad",
*    "value": "ddd"
*}
**/
func getHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := &model.Store{}

		// get POSTed data
		json.NewDecoder(r.Body).Decode(&store)
		c := make(chan error)

		// Run DB func
		go store.GetValue(c)

		// Handler errors
		if err := <-c; err != nil {
			json.NewEncoder(w).Encode(struct {
				Err string
			}{err.Error()})
			return
		}

		// Output values
		json.NewEncoder(w).Encode(store)
		return
	}
}
