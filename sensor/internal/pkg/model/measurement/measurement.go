package measurement

import (
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/formatter"
	"github.com/denislavPetkov/sensor/internal/pkg/printer"
)

type Measurement struct {
	MeasuredAt time.Time `yaml:"measuredAt" json:"measuredAt"`
	Value      string    `yanl:"value" json:"value"`
	SensorId   string    `yaml:"sensorId" json:"sensorId"`
	DeviceId   string    `yaml:"deviceId" json:"deviceId"`
}

func CreateMeasurement(timeStamp time.Time, value string, sensorId string, deviceId string) Measurement {

	measurement := Measurement{
		MeasuredAt: timeStamp,
		Value:      value,
		SensorId:   sensorId,
		DeviceId:   deviceId,
	}

	return measurement
}

func (m Measurement) PrintMeasurement(format string, webUrl string) error {

	measurement, err := m.formatMeasurement(format)
	if err != nil {
		return err
	}
	printer.Println(measurement)

	if webUrl == "" {
		return nil
	}

	if format == string(formatter.Yaml) {
		measurement, err = m.formatMeasurement(string(formatter.Json))
		if err != nil {
			return err
		}
	}
	printer.CallWebHookUrl(measurement, webUrl)

	return nil
}

func (m Measurement) formatMeasurement(selectedFormat string) (string, error) {

	formattedMeasurement, err := formatter.FormatMeasurement(m, selectedFormat)
	if err != nil {
		return "", err
	}

	return formattedMeasurement, nil
}
