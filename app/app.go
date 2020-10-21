package app

import (
	"log"
	"net/http"

	configuration "github.com/jgleon/topsecret/Configuration"
	routers "github.com/jgleon/topsecret/localization"
)

//Start Inicializa la configuracion del app
func Start() {
	app := configuration.GetInstace()
	config := app.GetConfiguration()

	router := routers.ConfigureRoutes()
	log.Fatal(http.ListenAndServe(":"+config.ApplicationInfo.Port, router))
}
