package main

import (
	"log"
	"net/http"

	configuration "github.com/jgleon/topsecret/Configuration"
	controller "github.com/jgleon/topsecret/Controllers"

	_ "github.com/jgleon/topsecret/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Localization API
// @version 1.0
// @description Manage satellite messages and calculate the location of the spacecraft
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := configuration.GetInstace()
	config := app.GetConfiguration()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/topsecret", controller.AddTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", controller.AddTopSecretSplit).Methods("POST")
	router.HandleFunc("/topsecret_split", controller.GetTopSecretSplit).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":"+config.ApplicationInfo.Port, router))
}
