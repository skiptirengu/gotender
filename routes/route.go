package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginRequest{}.HandleFunc).Methods("POST")
	r.HandleFunc("/videos/{query}", VideoSearchRequest{}.HandleFunc).Methods("GET")
}
