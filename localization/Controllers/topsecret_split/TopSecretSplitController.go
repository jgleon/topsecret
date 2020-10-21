package topsecret_split

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	helper "github.com/jgleon/topsecret/localization/Helper"
	models "github.com/jgleon/topsecret/localization/Models"
	repository "github.com/jgleon/topsecret/localization/Repository"
	services "github.com/jgleon/topsecret/localization/Services"
)

type Controller struct {
	LocalizationService services.IReadLocationServices
}

func NewTopSecretController() Controller {
	return Controller{
		LocalizationService: services.NewLocationService(),
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
	satellities := []models.Satellite{}
	_ = json.NewDecoder(req.Body).Decode(&infoSatellite)

	satellite.Distance = infoSatellite.Distance
	satellite.Message = infoSatellite.Message
	satellite.Name = strings.ToLower(params["satellite_name"])

	if !helper.ValidateSatelliteNames(append(satellities, satellite)) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repo := repository.GetInstace()
	repo.SetSatellite(satellite)

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
	repo := repository.GetInstace()
	satellites := repo.GetSatellites()

	location := contr.LocalizationService.ReadLocation(satellites)

	if location.Message == "" || (location.Position.X == float32(0) && location.Position.Y == float32(0)) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(location)
	}
}
