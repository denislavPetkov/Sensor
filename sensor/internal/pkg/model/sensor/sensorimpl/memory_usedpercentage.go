package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type memoryUsedPercentageSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewMemoryUsedPercentageSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &memoryUsedPercentageSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (m *memoryUsedPercentageSensor) GetMeasurement() (*measurement.Measurement, error) {
	usedPercentMemory, err := m.getUsedMemoryPercentage()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, usedPercentMemory, m.sensorMetadata.Id, m.deviceId)

	return &measurement, nil
}

func (m *memoryUsedPercentageSensor) getUsedMemoryPercentage() (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get used percent memory")
	}

	usedPercentMemory := fmt.Sprintf("%v", memory.UsedPercent)

	return usedPercentMemory, nil
}

func (m *memoryUsedPercentageSensor) GetName() string {
	return m.sensorMetadata.Name
}

func (m *memoryUsedPercentageSensor) GetSensorGroups() []sensor.SensorGroup {
	return m.sensorMetadata.SensorGroups
}

func (m *memoryUsedPercentageSensor) SetSensorName(sensorName string) {
	m.sensorMetadata.Name = sensorName
}

func (m *memoryUsedPercentageSensor) SetDeviceId(deviceId string) {
	m.deviceId = deviceId
}

func (m *memoryUsedPercentageSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	m.sensorMetadata = sensorMetaData
}
