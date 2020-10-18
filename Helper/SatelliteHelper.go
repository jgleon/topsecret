package helper

import (
	"strings"

	configuration "github.com/jgleon/topsecret/Configuration"
	models "github.com/jgleon/topsecret/Models"
)

func ValidateSatelliteNames(satellites []models.Satellite) bool {
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

func FormatSateliteName(satellites []models.Satellite) []models.Satellite {
	if len(satellites) > 0 {
		for index, _ := range satellites {
			satellites[index].Name = strings.ToLower(satellites[index].Name)
		}
	}

	return satellites
}