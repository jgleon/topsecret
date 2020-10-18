package models

type InfoSatellite struct {
	Distance float32  `json:"Distance" example:"100"`
	Message  []string `json:"Message"`
}

type Satellite struct {
	Name     string   `json:"Name" example:"kenobi"`
	Distance float32  `json:"Distance" example:"100"`
	Message  []string `json:"Message"`
}

type Satellites struct {
	Satellites []Satellite `json:"Satellites"`
}
