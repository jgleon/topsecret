package main

import (
	"github.com/jgleon/topsecret/app"
	_ "github.com/jgleon/topsecret/docs"
)

// @title Localization API
// @version 1.0
// @description Manage satellite messages and calculate the location of the spacecraft
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app.Start()
}
