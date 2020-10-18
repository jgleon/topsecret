package controller

import (
	"encoding/json"
	"net/http"

	models "github.com/jgleon/topsecret/Models"
	services "github.com/jgleon/topsecret/Services"
)

func AddTopSecret(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var satellites []models.Satellite

	_ = json.NewDecoder(req.Body).Decode(&satellites)

	location := services.ReadLocation(satellites)

	if location.Message == "" || (location.Position.X == float32(0) && location.Position.Y == float32(0)) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(location)
	}
}