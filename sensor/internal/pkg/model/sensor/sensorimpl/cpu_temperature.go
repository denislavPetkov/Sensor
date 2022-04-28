package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/host"
	log "github.com/sirupsen/logrus"
)

type cpuTemperatureSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewCpuTemperatureSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &cpuTemperatureSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (c *cpuTemperatureSensor) GetMeasurement() (*measurement.Measurement, error) {

	temperature, err := c.getTemperature()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, temperature, c.sensorMetadata.Id, c.deviceId)

	return &measurement, nil

}

const temperatureSensor = "TC0P"

func (c *cpuTemperatureSensor) getTemperature() (string, error) {
	var temperature string = "0"

	data, err := host.SensorsTemperatures()
	if err != nil {
		log.Error("Failed to retrieve host's temperature")
		return "", fmt.Errorf("Failed to read CPU's temperature")
	}

	for _, value := range data {
		if value.SensorKey != temperatureSensor {
			continue
		}
		temperature = fmt.Sprintf("%v", value.Temperature)
	}

	return temperature, nil
}

func (c *cpuTemperatureSensor) GetName() string {
	return c.sensorMetadata.Name
}

func (c *cpuTemperatureSensor) GetSensorGroups() []sensor.SensorGroup {
	return c.sensorMetadata.SensorGroups
}

func (c *cpuTemperatureSensor) SetSensorName(sensorName string) {
	c.sensorMetadata.Name = sensorName
}

func (c *cpuTemperatureSensor) SetDeviceId(deviceId string) {
	c.deviceId = deviceId
}

func (c *cpuTemperatureSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	c.sensorMetadata = sensorMetaData
}
