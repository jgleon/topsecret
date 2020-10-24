package helper

import (
	"strings"

	configuration "github.com/jgleon/topsecret/configuration"
	models "github.com/jgleon/topsecret/localization/models"
)

//IHelperFunction interfaz para funciones del helper
type IHelperFunction interface {
	ValidateSatelliteNames(satellites []models.Satellite) bool
	FormatSateliteName(satellites []models.Satellite) []models.Satellite
}

//HelperFunction estructura del helper
type HelperFunction struct {
}

//NewHelperFunction crea una instancia del helper
func NewHelperFunction() IHelperFunction {
	return &HelperFunction{}
}

//ValidateSatelliteNames Valida que los nombres de los satelites existan
func (h *HelperFunction) ValidateSatelliteNames(satellites []models.Satellite) bool {
	app := configuration.GetInstace()
	config := app.GetConfiguration()

	if len(satellites) > 0 {
		for _, valueSat := range satellites {
			found := false
			for _, valueConf := range config.SatellitesLocation {
				if strings.ToLower(valueSat.Name) == strings.ToLower(valueConf.Name) {
					found = true
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}

//FormatSateliteName formato a los nombres de los satelites
func (h *HelperFunction) FormatSateliteName(satellites []models.Satellite) []models.Satellite {
	if len(satellites) > 0 {
		for index := range satellites {
			satellites[index].Name = strings.ToLower(satellites[index].Name)
		}
	}

	return satellites
}
