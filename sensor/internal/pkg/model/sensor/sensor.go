package sensor

import (
	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
)

type Sensor struct {
	Id           string        `yaml:"id"`
	Name         string        `yaml:"name"`
	Description  string        `yaml:"description"`
	Unit         string        `yaml:"unit"`
	SensorGroups []SensorGroup `yaml:"sensorGroups"`
}

type SensorGroup string

type SensorInterface interface {
	GetMeasurement() (*measurement.Measurement, error)
	GetName() string
	GetSensorGroups() []SensorGroup
	SetSensorName(sensorName string)
	SetDeviceId(deviceId string)
	SetSensorMetaData(sensorMetaData Sensor)
}

func (s Sensor) GetName() string {
	return s.Name
}
