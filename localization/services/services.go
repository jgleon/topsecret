package services

import (
	"sort"
	"strings"

	geolocation "github.com/jgleon/topsecret/localization/Geolocation"
	helper "github.com/jgleon/topsecret/localization/Helper"
	models "github.com/jgleon/topsecret/localization/Models"
	repository "github.com/jgleon/topsecret/localization/Repository"
)

//ILocationServices interfaz del servicio
type ILocationServices interface {
	ReadLocation(satellites []models.Satellite) (location models.Location)
	SaveMessage(satellite models.Satellite) bool
	GetSatellites() (location models.Location)
}

//LocationServices funcion de localizacion
type LocationServices struct {
	Geolocation geolocation.IGetLocationFunc
	HelperFunc  helper.IHelperFunction
}

//NewLocationService crea una instancia del servicio
func NewLocationService() ILocationServices {
	return &LocationServices{
		Geolocation: geolocation.NewGetLocation(),
		HelperFunc:  helper.NewHelperFunction(),
	}
}

//ReadLocation obtiene la ubicación y el mensaje
func (serv *LocationServices) ReadLocation(satellites []models.Satellite) (location models.Location) {
	var distances []float32
	var messages [][]string
	location = models.Location{}

	if len(satellites) > 2 && serv.HelperFunc.ValidateSatelliteNames(satellites) {
		satellites = serv.HelperFunc.FormatSateliteName(satellites)

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

//SaveMessage guarda un mensaje de un satelite
func (serv *LocationServices) SaveMessage(satellite models.Satellite) bool {
	resultSave := false
	satellities := []models.Satellite{}

	if !serv.HelperFunc.ValidateSatelliteNames(append(satellities, satellite)) {
		return resultSave
	}

	repo := repository.GetInstace()
	repo.SetSatellite(satellite)

	resultSave = true

	return resultSave
}

// GetSatellites obtiene la ubicacion de la nave y el mensaje
func (serv *LocationServices) GetSatellites() (location models.Location) {
	location = models.Location{}

	repo := repository.GetInstace()
	satellites := repo.GetSatellites()
	location = serv.ReadLocation(satellites)

	return location
}
