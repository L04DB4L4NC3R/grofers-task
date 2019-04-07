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

/**
*@api {get} / get all values
*@apiName get all values
*@apiGroup all
*@apiParam {string} key key of the value
*@apiParamExample {json} request-example
*{
*    "key": "dsad"
*}
*@apiParamExample {json} response-example
*
*[
*    {
*        "key": "adsad",
*        "value": "adjhsbdbaa"
*    },
*    {
*        "key": "angad",
*        "value": "sharmazasd"
*    },
*    {
*        "key": "angadsd",
*        "value": "sharma"
*    },
*    {
*        "key": "da",
*        "value": "angad"
*    },
*    {
*        "key": "dasdsaasddasda",
*        "value": "angad"
*    },
*    {
*        "key": "dasdsaasddasdasad",
*        "value": "angad"
*    },
*    {
*        "key": "dsaasddasda",
*        "value": "angad"
*    },
*    {
*        "key": "dsadasda",
*        "value": "angad"
*    },
*    {
*        "key": "hip",
*        "value": "hop"
*    },
*    {
*        "key": "k2",
*        "value": "v2"
*    },
*    {
*        "key": "my",
*        "value": "name"
*    },
*    {
*        "key": "name",
*        "value": "angad"
*    },
*    {
*        "key": "name1",
*        "value": "angad"
*    },
*    {
*        "key": "name3",
*        "value": "angad"
*    },
*    {
*        "key": "name4",
*        "value": "angad"
*    },
*    {
*        "key": "name5",
*        "value": "angad"
*    },
*    {
*        "key": "name7",
*        "value": "angad"
*    },
*    {
*        "key": "new",
*        "value": "name"
*    },
*    {
*        "key": "ping",
*        "value": "pong"
*    }
*]
**/
func getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data, err := model.GetAll()

		// Handler errors
		if err != nil {
			json.NewEncoder(w).Encode(struct {
				Err string
			}{err.Error()})
			return
		}

		// Output values
		json.NewEncoder(w).Encode(data)
		return
	}
}
