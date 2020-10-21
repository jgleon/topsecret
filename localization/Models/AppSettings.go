package models

type SatellitesLocation struct {
	Name string
	X    float32
	Y    float32
}

type ApplicationInfo struct {
	Port string
}

type Settings struct {
	ApplicationInfo    ApplicationInfo
	SatellitesLocation []SatellitesLocation
}