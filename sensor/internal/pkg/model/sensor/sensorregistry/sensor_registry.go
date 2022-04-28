package sensorregistry

import (
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	log "github.com/sirupsen/logrus"
)

type SensorRegistry struct {
	registeredSensors []sensor.SensorInterface
}

func (s *SensorRegistry) GetSensorsByGroup(desiredSensorGroup string) []sensor.SensorInterface {
	var sensors []sensor.SensorInterface = make([]sensor.SensorInterface, 0)

	for _, currentSensor := range s.registeredSensors {
		sensorGroups := currentSensor.GetSensorGroups()
		for _, currentSensorGroup := range sensorGroups {
			if currentSensorGroup == sensor.SensorGroup(desiredSensorGroup) {
				sensors = append(sensors, currentSensor)
			}
		}
	}

	return sensors
}

func (s *SensorRegistry) RegisterSensors(sensors []sensor.SensorInterface) {
	s.registeredSensors = sensors
	log.Infof("Registering sensor objects with implementation into the registry")
}

func NewSensorRegistry() SensorRegistryInterface {

	sensorRegistry := &SensorRegistry{}

	return sensorRegistry
}

func (s *SensorRegistry) GetSensorByName(sensorName string) sensor.SensorInterface {
	for _, sensor := range s.registeredSensors {
		if currentSensorName := sensor.GetName(); currentSensorName == sensorName {
			return sensor
		}
	}
	return nil
}

type SensorRegistryInterface interface {
	GetSensorByName(sensorName string) sensor.SensorInterface
	RegisterSensors(sensors []sensor.SensorInterface)
	GetSensorsByGroup(desiredSensorGroup string) []sensor.SensorInterface
}
