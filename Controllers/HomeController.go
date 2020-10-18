package controller

import (
	"net/http"
)

func HomeController(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Geolocation by the trilateration algorithm"))
}
