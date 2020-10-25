package configuration

import (
	"os"
	"sync"

	models "github.com/jgleon/topsecret/localization/models"
	"github.com/tkanos/gonfig"
)

var (
	app  *application
	once sync.Once
)

type application struct {
	Settings models.Settings
}

//GetInstace crea una instancia del singleton
func GetInstace() *application {
	once.Do(func() {
		app = &application{}
		fileName := getPathFile()
		gonfig.GetConf(fileName, &app.Settings)
	})

	return app
}

func (app *application) GetConfiguration() models.Settings {
	return app.Settings
}

func getPathFile() string {
	filePath := "./Settings/appsettings.json"

	if _, err := os.Stat(filePath); err != nil {
		filePath = "../Settings/appsettings.json"

		if _, err := os.Stat(filePath); err != nil {
			filePath = "../../Settings/appsettings.json"

			if _, err := os.Stat(filePath); err != nil {
				filePath = "../../../Settings/appsettings.json"
			}
		}
	}

	return filePath
}
