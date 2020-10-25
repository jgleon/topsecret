package topsecret

import (
	models "github.com/jgleon/topsecret/localization/models"
)

func getBodyMock() models.Satellites {
	satellites := models.Satellites{}
	arraySatellites := []models.Satellite{}
	satellite1 := models.Satellite{}
	satellite1.Name = "kenobi"
	satellite1.Distance = 900
	satellite1.Message = []string{"Esto", "", "un", "", ""}
	arraySatellites = append(arraySatellites, satellite1)

	satellite2 := models.Satellite{}
	satellite2.Name = "skywalker"
	satellite2.Distance = 415
	satellite2.Message = []string{"", "es", "", "mensaje", ""}
	arraySatellites = append(arraySatellites, satellite2)

	satellite3 := models.Satellite{}
	satellite3.Name = "sato"
	satellite3.Distance = 342
	satellite3.Message = []string{"esto", "", "", "", "secreto"}
	arraySatellites = append(arraySatellites, satellite3)
	satellites.Satellites = arraySatellites
	return satellites

}

func expectedCreateUserBody() models.Location {
	expected := models.Location{}
	expected.Message = "Esto es un mensaje secreto"
	expected.Position.X = 275.18063
	expected.Position.Y = 187.79124
	return expected
}
