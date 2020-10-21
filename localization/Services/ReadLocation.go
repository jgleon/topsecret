package services

import (
	"sort"
	"strings"

	geolocation "github.com/jgleon/topsecret/localization/Geolocation"
	helper "github.com/jgleon/topsecret/localization/Helper"
	models "github.com/jgleon/topsecret/localization/Models"
)

//IReadLocationServices interfaz del servicio
type IReadLocationServices interface {
	ReadLocation(satellites []models.Satellite) (location models.Location)
}

//ReadLocationServices funcion de localizacion
type ReadLocationServices struct {
	Geolocation geolocation.IGetLocationFunc
}

//NewLocationService crea una instancia del servicio
func NewLocationService() IReadLocationServices {
	return &ReadLocationServices{
		Geolocation: geolocation.NewGetLocation(),
	}
}

//ReadLocation obtiene la ubicación y el mensaje
func (serv *ReadLocationServices) ReadLocation(satellites []models.Satellite) (location models.Location) {
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

		location.Position.X, location.Position.Y = serv.Geolocation.GetLocation(distances)
		location.Message = serv.Geolocation.GetMessage(messages)
	}
	return location
}
