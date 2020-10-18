package services

import (
	"sort"
	"strings"

	geolocation "github.com/jgleon/topsecret/Geolocation"
	helper "github.com/jgleon/topsecret/Helper"
	models "github.com/jgleon/topsecret/Models"
)

func ReadLocation(satellites []models.Satellite) (location models.Location) {
	var distances []float32
	var messages [][]string
	location = models.Location{}

	if len(satellites) > 2 && helper.ValidateSatelliteNames(satellites) {
		satellites = helper.FormatSateliteName(satellites)

		sort.SliceStable(satellites, func(i, j int) bool {
			return strings.ToLower(satellites[i].Name) < strings.ToLower(satellites[j].Name)
		})

		for _, value := range satellites {
			distances = append(distances, value.Distance)
			messages = append(messages, value.Message)
		}

		location.Position.X, location.Position.Y = geolocation.GetLocation(distances)
		location.Message = geolocation.GetMessage(messages)
	}
	return location
}