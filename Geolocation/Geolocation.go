package geolocation

import (
	"math"
	"sort"
	"strings"

	configuration "github.com/jgleon/topsecret/Configuration"
)

func GetLocation(distances []float32) (x, y float32) {
	var _x float32 = float32(0)
	var _y float32 = float32(0)

	if len(distances) > 2 {
		app := configuration.GetInstace()
		config := app.GetConfiguration()
		satellitesLocation := config.SatellitesLocation

		sort.SliceStable(satellitesLocation, func(i, j int) bool {
			return satellitesLocation[i].Name < satellitesLocation[j].Name
		})

		x1 := float64(satellitesLocation[0].X)
		x2 := float64(satellitesLocation[1].X)
		x3 := float64(satellitesLocation[2].X)
		y1 := float64(satellitesLocation[0].Y)
		y2 := float64(satellitesLocation[1].Y)
		y3 := float64(satellitesLocation[2].Y)
		r1 := float64(distances[0])
		r2 := float64(distances[1])
		r3 := float64(distances[2])

		A := 2*x2 - 2*x1
		B := 2*y2 - 2*y1
		C := math.Pow(r1, 2) - math.Pow(r2, 2) - math.Pow(x1, 2) + math.Pow(x2, 2) - math.Pow(y1, 2) + math.Pow(y2, 2)
		D := 2*x3 - 2*x2
		E := 2*y3 - 2*y2
		F := math.Pow(r2, 2) - math.Pow(r3, 2) - math.Pow(x2, 2) + math.Pow(x3, 2) - math.Pow(y2, 2) + math.Pow(y3, 2)

		if (E*A-B*D) != float64(0) && (B*D-A*E) != float64(0) {
			_x = float32((C*E - F*B) / (E*A - B*D))
			_y = float32((C*D - A*F) / (B*D - A*E))
		}
	}

	return _x, _y
}

func GetMessage(messages [][]string) (message string) {
	var _message string = ""

	if len(messages) > 2 {
		messages = arrayAutocompleted(messages)
		var arrayLen int = len(messages[0])
		_messages := make([]string, arrayLen)

		for _, valueRow := range messages {
			for indexColumn, valueColumn := range valueRow {
				if _messages[indexColumn] == "" {
					_messages[indexColumn] = valueColumn
				} else if _messages[indexColumn] != "" && strings.ToLower(_messages[indexColumn]) != strings.ToLower(valueColumn) && valueColumn != "" {
					_messages[indexColumn] = ""
				}
			}
		}

		_message = strings.Join(_messages, " ")
	}

	return _message
}

func arrayAutocompleted(messages [][]string) (_messages [][]string) {
	var _size int = 0

	for i := 0; i < 2; i++ {
		for _, valueRow := range messages {
			_sizeRow := len(valueRow)

			if i == 0 {
				if _sizeRow > _size {
					_size = _sizeRow
				}
			} else {
				if _sizeRow < _size {
					missingColumns := make([]string, _size-_sizeRow)
					rowCompleted := append(valueRow, missingColumns...)
					_messages = append(_messages, rowCompleted)
				} else {
					_messages = append(_messages, valueRow)
				}
			}
		}
	}

	return _messages
}