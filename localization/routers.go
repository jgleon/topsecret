package localization

import (
	"github.com/gorilla/mux"
	topsecret_controller "github.com/jgleon/topsecret/localization/Controllers/topsecret"
	topsecret_split_controller "github.com/jgleon/topsecret/localization/Controllers/topsecret_split"
	httpSwagger "github.com/swaggo/http-swagger"
)

//ConfigureRoutes Configura las rutas del API
func ConfigureRoutes() *mux.Router {
	topSecretController := topsecret_controller.NewTopSecretController()
	topSecretSplitController := topsecret_split_controller.NewTopSecretController()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/topsecret", topSecretController.AddTopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", topSecretSplitController.AddTopSecretSplit).Methods("POST")
	router.HandleFunc("/topsecret_split", topSecretSplitController.GetTopSecretSplit).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return router
}
