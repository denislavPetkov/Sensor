package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
)

type gpuMemorySensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewGpuMemorySensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &gpuMemorySensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (g *gpuMemorySensor) GetMeasurement() (*measurement.Measurement, error) {
	gpuMemory := fmt.Sprintf("%v", 6)
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, gpuMemory, g.sensorMetadata.Id, g.deviceId)

	return &measurement, nil
}

func (g *gpuMemorySensor) GetName() string {
	return g.sensorMetadata.Name
}

func (g *gpuMemorySensor) GetSensorGroups() []sensor.SensorGroup {
	return g.sensorMetadata.SensorGroups
}

func (g *gpuMemorySensor) SetSensorName(sensorName string) {
	g.sensorMetadata.Name = sensorName
}

func (g *gpuMemorySensor) SetDeviceId(deviceId string) {
	g.deviceId = deviceId
}

func (g *gpuMemorySensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	g.sensorMetadata = sensorMetaData
}
