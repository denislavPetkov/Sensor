package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type memoryAvailableSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewMemoryAvailableSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &memoryAvailableSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (m *memoryAvailableSensor) GetMeasurement() (*measurement.Measurement, error) {
	availableMemory, err := m.getAvailableMemory()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, availableMemory, m.sensorMetadata.Id, m.deviceId)

	return &measurement, nil
}

func (m *memoryAvailableSensor) getAvailableMemory() (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get available memory")
	}

	availableMemory := fmt.Sprintf("%v", memory.Available)

	return availableMemory, nil
}

func (m *memoryAvailableSensor) GetName() string {
	return m.sensorMetadata.Name
}

func (m *memoryAvailableSensor) GetSensorGroups() []sensor.SensorGroup {
	return m.sensorMetadata.SensorGroups
}

func (m *memoryAvailableSensor) SetSensorName(sensorName string) {
	m.sensorMetadata.Name = sensorName
}

func (m *memoryAvailableSensor) SetDeviceId(deviceId string) {
	m.deviceId = deviceId
}

func (m *memoryAvailableSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	m.sensorMetadata = sensorMetaData
}
