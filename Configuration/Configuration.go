package configuration

import (
	"sync"

	models "github.com/jgleon/topsecret/Models"
	"github.com/tkanos/gonfig"
)

var (
	app  *application
	once sync.Once
)

type application struct {
	Settings models.Settings
}

func GetInstace() *application {
	once.Do(func() {
		app = &application{}

		fileName := "./Settings/appsettings.json"
		gonfig.GetConf(fileName, &app.Settings)
	})

	return app
}

func (app *application) GetConfiguration() models.Settings {
	return app.Settings
}
