package topsecretsplit

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	models "github.com/jgleon/topsecret/localization/models"
	services "github.com/jgleon/topsecret/localization/services"
)

//Controller topsecret split
type Controller struct {
	LocalizationService services.ILocationServices
}

//NewTopSecretController crea una instancia del controller
func NewTopSecretController(service services.ILocationServices) Controller {
	return Controller{
		LocalizationService: service,
	}
}

// AddTopSecretSplit godoc
// @Summary Records a satellite message in memory
// @Description Records a satellite message in memory to calculate the location of the spacecraft later
// @Tags topsecret_split
// @Accept  json
// @Produce  json
// @Param satellite_name path string true "satellite name"
// @Param infoSatellite body models.InfoSatellite true "infomation message satellite"
// @Success 202 {string} string ""
// @Failure 404
// @Router /topsecret_split/{satellite_name} [post]
func (contr *Controller) AddTopSecretSplit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var infoSatellite models.InfoSatellite
	satellite := models.Satellite{}

	_ = json.NewDecoder(req.Body).Decode(&infoSatellite)

	satellite.Distance = infoSatellite.Distance
	satellite.Message = infoSatellite.Message
	satellite.Name = strings.ToLower(params["satellite_name"])

	resultSave := contr.LocalizationService.SaveMessage(satellite)

	if !resultSave {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// GetTopSecretSplit godoc
// @Summary Determine the location of the spaceship
// @Description Determine the location of the spacecraft using the data sent
// @Tags topsecret_split
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Location
// @Failure 404
// @Router /topsecret_split [get]
func (contr *Controller) GetTopSecretSplit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	location := contr.LocalizationService.GetSatellites()

	if location.Message == "" || (location.Position.X == float32(0) && location.Position.Y == float32(0)) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(location)
	}
}
