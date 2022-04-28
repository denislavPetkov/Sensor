package sensorimpl

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
)

type memoryTotalSensor struct {
	sensorMetadata sensor.Sensor
	deviceId       string
}

func NewMemoryTotalSensor(sensorMetaData sensor.Sensor, deviceId string) sensor.SensorInterface {
	sensor := &memoryTotalSensor{
		sensorMetadata: sensorMetaData,
		deviceId:       deviceId,
	}
	return sensor
}

func (m *memoryTotalSensor) GetMeasurement() (*measurement.Measurement, error) {
	totalMemory, err := m.getTotalMemory()
	if err != nil {
		return nil, err
	}
	timeStamp := time.Now().Truncate(time.Second)
	measurement := measurement.CreateMeasurement(timeStamp, totalMemory, m.sensorMetadata.Id, m.deviceId)

	return &measurement, nil
}

func (m *memoryTotalSensor) getTotalMemory() (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Failed to get total memory")
	}
	totalMemoryGB := memory.Total >> 30

	totalMemory := fmt.Sprintf("%v", totalMemoryGB)

	return totalMemory, nil
}

func (m *memoryTotalSensor) GetName() string {
	return m.sensorMetadata.Name
}

func (m *memoryTotalSensor) GetSensorGroups() []sensor.SensorGroup {
	return m.sensorMetadata.SensorGroups
}

func (m *memoryTotalSensor) SetSensorName(sensorName string) {
	m.sensorMetadata.Name = sensorName
}

func (m *memoryTotalSensor) SetDeviceId(deviceId string) {
	m.deviceId = deviceId
}

func (m *memoryTotalSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	m.sensorMetadata = sensorMetaData
}
