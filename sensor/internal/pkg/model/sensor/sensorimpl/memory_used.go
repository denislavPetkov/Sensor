package sensorimpl

import (
	"fmt"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
)

type memoryUsedSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewMemoryUsedSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &memoryUsedSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (m *memoryUsedSensor) GetMeasurement() (*measurement.Measurement, error) {
	usedMemory, err := m.getUsedMemory()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, usedMemory, m.sensorMetadata.Id, m.deviceId)

	return &measurement, nil
}

func (m *memoryUsedSensor) getUsedMemory() (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get used memory")
	}

	usedMemory := fmt.Sprintf("%v", memory.Used)

	return usedMemory, nil
}

func (m *memoryUsedSensor) GetName() string {
	return m.sensorMetadata.Name
}

func (m *memoryUsedSensor) GetSensorGroups() []sensor.SensorGroup {
	return m.sensorMetadata.SensorGroups
}

func (m *memoryUsedSensor) SetSensorName(sensorName string) {
	m.sensorMetadata.Name = sensorName
}

func (m *memoryUsedSensor) SetDeviceId(deviceId string) {
	m.deviceId = deviceId
}

func (m *memoryUsedSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	m.sensorMetadata = sensorMetaData
}
