package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/context"
	"github.com/skiptirengu/gotender/middlewares"
	"encoding/json"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginRequest{}.HandleFunc).Methods("POST")

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		val, _ := json.Marshal(context.Get(request, middlewares.UserContext))
		writer.Write(val)
	})
}
