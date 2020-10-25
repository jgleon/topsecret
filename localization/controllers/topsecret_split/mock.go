package topsecretsplit

import (
	models "github.com/jgleon/topsecret/localization/models"
)

func getBodySplitMock(index int) (models.InfoSatellite, string) {
	infoSatellite := models.InfoSatellite{}
	satelliteName := ""

	if index == 0 {
		satelliteName = "kenobi"
		infoSatellite.Distance = 900
		infoSatellite.Message = []string{"Esto", "", "un", "", ""}
	}

	if index == 1 {
		satelliteName = "skywalker"
		infoSatellite.Distance = 415
		infoSatellite.Message = []string{"", "es", "", "mensaje", ""}
	}

	if index == 2 {
		satelliteName = "sato"
		infoSatellite.Distance = 342
		infoSatellite.Message = []string{"esto", "", "", "", "secreto"}
	}

	return infoSatellite, satelliteName
}

func expectedGetLocationSplitBody() models.Location {
	expected := models.Location{}
	expected.Message = "Esto es un mensaje secreto"
	expected.Position.X = 275.18063
	expected.Position.Y = 187.79124
	return expected
}
