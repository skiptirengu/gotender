package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginRequest{}.HandleFunc).Methods("POST")

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world"))
	})
}
