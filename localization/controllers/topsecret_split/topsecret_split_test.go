package topsecretsplit

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
	router.HandleFunc("/topsecret_split/{satellite_name}", ctrl.AddTopSecretSplit)
	router.HandleFunc("/topsecret_split", ctrl.GetTopSecretSplit)
	return router
}

func TestSaveMessageSatellite(t *testing.T) {
	route := Router()

	for x := 0; x < 3; x++ {
		body, satelliteName := getBodySplitMock(x)
		payload, _ := json.Marshal(body)

		path := "/topsecret_split/" + satelliteName
		req, _ := http.NewRequest("POST", path, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusAccepted, rr.Code)
	}
}

func TestCalculateLocationSplit(t *testing.T) {
	route := Router()

	for x := 0; x < 3; x++ {
		body, satelliteName := getBodySplitMock(x)
		payload, _ := json.Marshal(body)

		path := "/topsecret_split/" + satelliteName
		req, _ := http.NewRequest("POST", path, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		route.ServeHTTP(rr, req)
	}

	req, _ := http.NewRequest("GET", "/topsecret_split", nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)

	result := &models.Location{}
	json.Unmarshal(rr.Body.Bytes(), &result)

	expected := expectedGetLocationSplitBody()

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expected.Message, result.Message)
	assert.Equal(t, expected.Position.X, result.Position.X)
	assert.Equal(t, expected.Position.Y, result.Position.Y)
}
