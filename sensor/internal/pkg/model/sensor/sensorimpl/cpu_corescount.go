package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/cpu"
	log "github.com/sirupsen/logrus"
)

type cpuCoresCountSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewCpuCoresCountSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &cpuCoresCountSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (c *cpuCoresCountSensor) GetMeasurement() (*measurement.Measurement, error) {
	cpuCores, err := c.getCoresCount()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, cpuCores, c.sensorMetadata.Id, c.deviceId)

	return &measurement, nil
}

func (c *cpuCoresCountSensor) getCoresCount() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get CPU cores count")
	}

	cpuCores := fmt.Sprintf("%v", cpuInfo[0].Cores)

	return cpuCores, nil
}

func (c *cpuCoresCountSensor) GetName() string {
	return c.sensorMetadata.Name
}

func (c *cpuCoresCountSensor) GetSensorGroups() []sensor.SensorGroup {
	return c.sensorMetadata.SensorGroups
}

func (c *cpuCoresCountSensor) SetSensorName(sensorName string) {
	c.sensorMetadata.Name = sensorName
}

func (c *cpuCoresCountSensor) SetDeviceId(deviceId string) {
	c.deviceId = deviceId
}

func (c *cpuCoresCountSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	c.sensorMetadata = sensorMetaData
}
