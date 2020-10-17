package main

import (
	"net/http"

	configuration "github.com/jgleon/topsecret/Configuration"
	controller "github.com/jgleon/topsecret/Controllers"

	"github.com/gorilla/mux"
)

func main() {
	app := configuration.GetInstace()
	config := app.GetConfiguration()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.HomeController)
	router.HandleFunc("/topsecret", controller.AddTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", controller.AddTopSecretSplit).Methods("POST")
	router.HandleFunc("/topsecret_split", controller.GetTopSecretSplit).Methods("GET")

	http.ListenAndServe(":"+config.ApplicationInfo.Port, router)
}
