package topsecret

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"github.com/gorilla/mux"
// 	models "github.com/jgleon/topsecret/localization/models"
// 	"github.com/stretchr/testify/assert"
// )

// var Router *mux.Router

// type SuiteTopSecretController struct {
// 	topSecretServiceMock *ServiceMock
// 	controller           Controller
// }

// func TestTopSecret(t *testing.T) {
// 	topSecretServiceMock := &ServiceMock{}
// 	controller := Controller{TopSecretService: topSecretServiceMock}
// 	topSecretServiceMock.readLocation = func(satellites *[]models.Satellite) (location *models.Location) {
// 		location = &models.Location{
// 			Message: "Esto es un mensaje secreto",
// 		}
// 		location.Position.X = 274.7016
// 		location.Position.Y = 188.58919
// 		return location
// 	}

// 	responseRecorder, context := createTestContextAndRecorder(http.MethodPost, "/topsecret", getBodyMock())
// 	// nuevo router
// 	controller.AddTopSecret(context)
// 	response := responseRecorder.Result()
// 	bodyString := getStringBody(response)

// 	assert.Equal(t, http.StatusOK, response.StatusCode)
// 	assert.Equal(t, expectedCreateUserBody(), bodyString)
// }

// func expectedCreateUserBody() string {
// 	return fmt.Sprintf(`{"Position":{"x":274.7016,"y":188.58919},"message": "Esto es un mensaje secreto"}`)
// }
