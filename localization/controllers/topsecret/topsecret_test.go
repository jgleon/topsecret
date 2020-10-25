package topsecret

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	configuration "github.com/jgleon/topsecret/configuration"
	models "github.com/jgleon/topsecret/localization/models"
	services "github.com/jgleon/topsecret/localization/services"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	app := configuration.GetInstace()
	_ = app.GetConfiguration()

	router := mux.NewRouter()
	ctrl := NewTopSecretController(services.NewLocationService())
	router.HandleFunc("/topsecret", ctrl.AddTopSecret)
	return router
}

func TestCalculateLocation(t *testing.T) {
	payload, _ := json.Marshal(getBodyMock())

	req, _ := http.NewRequest("POST", "/topsecret", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	result := &models.Location{}
	json.Unmarshal(rr.Body.Bytes(), &result)

	expected := expectedCreateUserBody()

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected.Message, result.Message)
	assert.Equal(t, expected.Position.X, result.Position.X)
	assert.Equal(t, expected.Position.Y, result.Position.Y)
}
