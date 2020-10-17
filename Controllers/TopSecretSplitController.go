package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	helper "github.com/jgleon/topsecret/Helper"
	models "github.com/jgleon/topsecret/Models"
	repository "github.com/jgleon/topsecret/Repository"
	services "github.com/jgleon/topsecret/Services"
)

func AddTopSecretSplit(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var satellite models.Satellite
	var satellities []models.Satellite
	_ = json.NewDecoder(req.Body).Decode(&satellite)

	satellite.Name = strings.ToLower(params["satellite_name"])

	if !helper.ValidateSatelliteNames(append(satellities, satellite)) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	repo := repository.GetInstace()
	repo.SetSatellite(satellite)

	w.WriteHeader(http.StatusAccepted)
}

func GetTopSecretSplit(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	repo := repository.GetInstace()
	satellites := repo.GetSatellites()

	location := services.ReadLocation(satellites)

	if location.Message == "" || (location.Position.X == float32(0) && location.Position.Y == float32(0)) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(location)
	}
}
