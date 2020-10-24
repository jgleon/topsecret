package topsecret

import (
	"encoding/json"
	"net/http"

	models "github.com/jgleon/topsecret/localization/models"
	services "github.com/jgleon/topsecret/localization/services"
)

//Controller contiene los servicios del controller
type Controller struct {
	LocalizationService services.ILocationServices
}

//NewTopSecretController crea una instancia del controller
func NewTopSecretController(service services.ILocationServices) Controller {
	return Controller{
		LocalizationService: service,
	}
}

// AddTopSecret godoc
// @Summary Determine the location of the spaceship
// @Description Determine the location of the spacecraft using the data sent
// @Tags topsecret
// @Accept  json
// @Produce  json
// @Param satellites body models.Satellites true "information messages from satellites"
// @Success 200 {object} models.Location
// @Failure 404
// @Router /topsecret [post]
func (contr *Controller) AddTopSecret(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var satellites models.Satellites

	_ = json.NewDecoder(req.Body).Decode(&satellites)

	location := contr.LocalizationService.ReadLocation(satellites.Satellites)

	if location.Message == "" || (location.Position.X == float32(0) && location.Position.Y == float32(0)) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(location)
	}
}
