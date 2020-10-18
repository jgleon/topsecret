package models

type Location struct {
	Position struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
	}
	Message string `json:"message"`
}