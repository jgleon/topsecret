package models

type ErrorResponse struct {
	Code    string   `json:"Name"`
	Message []string `json:"Message"`
}