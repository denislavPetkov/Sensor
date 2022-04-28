package device

import (
	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorregistry"
	"github.com/denislavPetkov/sensor/internal/pkg/printer"
	log "github.com/sirupsen/logrus"
)

type Device struct {
	Id             string          `yaml:"id"`
	Name           string          `yaml:"name"`
	Description    string          `yaml:"description"`
	Sensors        []sensor.Sensor `yaml:"sensors"`
	SensorRegistry sensorregistry.SensorRegistryInterface
}

func (device *Device) GetSensors() []sensor.Sensor {
	return device.Sensors
}

func (device *Device) GetMeasurements(sensorGroups []string) ([]measurement.Measurement, error) {
	var availableSensors map[string]sensor.SensorInterface = make(map[string]sensor.SensorInterface)
	var measurements []measurement.Measurement = make([]measurement.Measurement, 0)
	var err error

	for _, sensorGroup := range sensorGroups {
		registeredSensors := device.SensorRegistry.GetSensorsByGroup(sensorGroup)
		for _, sensor := range registeredSensors {
			sensorName := sensor.GetName()
			availableSensors[sensorName] = sensor
		}
	}

	measurements, err = device.getMeasurements(availableSensors)
	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (device *Device) getMeasurements(sensors map[string]sensor.SensorInterface) ([]measurement.Measurement, error) {
	var measurements []measurement.Measurement = make([]measurement.Measurement, 0)

	for _, sensor := range sensors {
		measurement, err := sensor.GetMeasurement()
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, *measurement)
	}
	return measurements, nil
}

func (device *Device) SetRegistry(registry sensorregistry.SensorRegistryInterface) {
	device.SensorRegistry = registry
}

func (device *Device) GetId() string {
	return device.Id
}

func (device *Device) DoRegisteredSensorsExist(sensorGroups []string) bool {

	var doRegisteredSensorsExist bool = false

	for _, sensorGroup := range sensorGroups {
		registeredSensors := device.SensorRegistry.GetSensorsByGroup(sensorGroup)
		if len(registeredSensors) == 0 {
			log.Infof("Device with id '%v' has no sensors from '%v' group", device.Id, sensorGroup)
			printer.Printf("Device with id '%v' has no sensors from '%v' group\n", device.Id, sensorGroup)
		} else {
			doRegisteredSensorsExist = true
		}
	}
	return doRegisteredSensorsExist
}
