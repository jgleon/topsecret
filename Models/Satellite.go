package models

type Satellite struct {
	Name     string   `json:"Name"`
	Distance float32  `json:"Distance"`
	Message  []string `json:"Message"`
}