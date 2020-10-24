package topsecret

import (
	models "github.com/jgleon/topsecret/localization/models"
)

//ServiceMock mock del servicio
type ServiceMock struct {
	readLocation func(satellites []models.Satellite) (location models.Location)
	// saveMessage   func(satellite models.Satellite) bool
	// getSatellites func(location models.Location)
}

//var _ service.ILocationServices = (*ServiceMock)

//ReadLocation crea una funcion mock
func (s *ServiceMock) ReadLocation(satellites []models.Satellite) (location models.Location) {
	// if s.readLocation == nil {
	// 	return &models.Location{}
	// }
	return s.ReadLocation(satellites)
}

func getBodyMock() string {
	return `{
		"satellites": [
			{
				"Name": "kenobi",
				"Distance": 900,
				"Message": [
					"Esto",
					"",
					"un",
					"",
					""
				]
			},
			{
				"Name": "skywalker",
				"Distance": 415,
				"Message": [
					"",
					"es",
					"",
					"mensaje",
					""
				]
			},
			{
				"Name": "sato",
				"Distance": 342,
				"Message": [
					"esto",
					"",
					"",
					"",
					"secreto"
				]
			}
		]
	}`
}
