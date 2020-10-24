package localization

import (
	"github.com/gorilla/mux"
	topsecret_controller "github.com/jgleon/topsecret/localization/controllers/topsecret"
	topsecret_split_controller "github.com/jgleon/topsecret/localization/controllers/topsecret_split"
	"github.com/jgleon/topsecret/localization/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

//ConfigureRoutes Configura las rutas del API
func ConfigureRoutes() *mux.Router {
	service := services.NewLocationService()
	topSecretController := topsecret_controller.NewTopSecretController(service)
	topSecretSplitController := topsecret_split_controller.NewTopSecretController(service)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/topsecret", topSecretController.AddTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", topSecretSplitController.AddTopSecretSplit).Methods("POST")
	router.HandleFunc("/topsecret_split", topSecretSplitController.GetTopSecretSplit).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return router
}
