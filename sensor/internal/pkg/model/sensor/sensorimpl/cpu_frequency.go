package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

type cpuFrequencySensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewCpuFrequencySensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &cpuFrequencySensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (c *cpuFrequencySensor) GetMeasurement() (*measurement.Measurement, error) {
	cpuFrequencySensor, err := c.getFrequency()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, cpuFrequencySensor, c.sensorMetadata.Id, c.deviceId)

	return &measurement, nil
}

func (c *cpuFrequencySensor) getFrequency() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get CPU frequency")
	}

	cpuFrequency := fmt.Sprintf("%v", cpuInfo[0].Mhz)

	return cpuFrequency, nil
}

func (c *cpuFrequencySensor) GetName() string {
	return c.sensorMetadata.Name
}

func (c *cpuFrequencySensor) GetSensorGroups() []sensor.SensorGroup {
	return c.sensorMetadata.SensorGroups
}

func (c *cpuFrequencySensor) SetSensorName(sensorName string) {
	c.sensorMetadata.Name = sensorName
}

func (c *cpuFrequencySensor) SetDeviceId(deviceId string) {
	c.deviceId = deviceId
}

func (c *cpuFrequencySensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	c.sensorMetadata = sensorMetaData
}
