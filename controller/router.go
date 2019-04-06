package controller

import (
	"net/http"

	"github.com/rs/cors"
)

var taskID uint

func Startup() *http.Handler {
	m := http.NewServeMux()

	// handle CORS
	corsHandler := cors.Default().Handler(m)
	return &corsHandler
}
