package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

type cpuUsagePercentageSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewCpuUsagePercentageSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &cpuUsagePercentageSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (c *cpuUsagePercentageSensor) GetMeasurement() (*measurement.Measurement, error) {
	cpuUsagePercentageSensor, err := c.getUsagePercentage()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, cpuUsagePercentageSensor, c.sensorMetadata.Id, c.deviceId)

	return &measurement, nil
}

func (c *cpuUsagePercentageSensor) getUsagePercentage() (string, error) {
	usage, err := cpu.Percent(0, false)
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get CPU usage")
	}

	cpuUsage := fmt.Sprintf("%v", usage[0])

	return cpuUsage, nil
}

func (c *cpuUsagePercentageSensor) GetName() string {
	return c.sensorMetadata.Name
}

func (c *cpuUsagePercentageSensor) GetSensorGroups() []sensor.SensorGroup {
	return c.sensorMetadata.SensorGroups
}

func (c *cpuUsagePercentageSensor) SetSensorName(sensorName string) {
	c.sensorMetadata.Name = sensorName
}

func (c *cpuUsagePercentageSensor) SetDeviceId(deviceId string) {
	c.deviceId = deviceId
}

func (c *cpuUsagePercentageSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	c.sensorMetadata = sensorMetaData
}
