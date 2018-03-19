package main

import (
	"github.com/skiptirengu/gotender/models"
	"github.com/skiptirengu/gotender/middlewares"
	"github.com/skiptirengu/gotender/routes"
	"github.com/gorilla/mux"

	"net/http"
	"log"
	"fmt"
	"github.com/skiptirengu/gotender/config"
)

func main() {
	// Init database
	if _, err := models.Migrate(); err != nil {
		panic(err.Error())
	}

	r := mux.NewRouter()
	middlewares.RegisterMiddlewares(r)
	routes.RegisterRoutes(r)

	port := fmt.Sprintf(":%d", config.Get().ApiPort)
	log.Fatal(http.ListenAndServe(port, r))
}
