package controller

import (
	"net/http"

	"github.com/rs/cors"
)

func Startup() *http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/api/v1/get", getHandler())
	m.HandleFunc("/api/v1/put", putHandler())
	m.HandleFunc("/api/v1/update", updateHandler())
	m.HandleFunc("/api/v1/delete", deleteHandler())
	m.HandleFunc("/api/v1/", getAll())
	corsMux := cors.Default().Handler(m)
	return &corsMux
}
