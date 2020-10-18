package repository

import (
	"sync"

	models "github.com/jgleon/topsecret/Models"
)

var (
	_sat *satellites
	once sync.Once
)

func GetInstace() *satellites {
	once.Do(func() {
		_sat = &satellites{}
	})

	return _sat
}

type satellites struct {
	list []models.Satellite
	mux  sync.Mutex
}

func (_sat *satellites) SetSatellite(sat models.Satellite) {
	_sat.mux.Lock()
	index := _sat.searchSatellite(sat.Name)

	if index < 0 {
		_sat.list = append(_sat.list, sat)
	} else {
		_sat.list[index].Message = sat.Message
		_sat.list[index].Distance = sat.Distance
	}
	_sat.mux.Unlock()
}

func (_sat *satellites) GetSatellites() []models.Satellite {
	_sat.mux.Lock()
	_sat.mux.Unlock()
	return _sat.list
}

func (_sat *satellites) searchSatellite(name string) (index int) {
	index = -1

	for i := range _sat.list {
		if _sat.list[i].Name == name {
			index = i
		}
	}
	return index
}