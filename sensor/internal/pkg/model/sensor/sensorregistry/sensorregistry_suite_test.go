package sensorregistry

import (
	"testing"

	_ "embed"

	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSensorregistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sensorregistry Suite")
}

const (
	deviceId    = "1"
	sensorId    = "1"
	sensorName  = "cpuTemperatureSensor"
	sensorGroup = "CPU_TEMP"
)

var (
	expectedMeasurement = measurement.Measurement{
		SensorId: sensorId,
		DeviceId: deviceId,
	}
)

type MockSensor struct {
	sensor.Sensor
	deviceId string
	measurement.Measurement
}

func (m *MockSensor) GetMeasurement() (*measurement.Measurement, error) {
	return &expectedMeasurement, nil
}

func (m *MockSensor) GetName() string {
	return m.Sensor.Name
}

func (m *MockSensor) GetSensorGroups() []sensor.SensorGroup {
	return m.Sensor.SensorGroups
}

func (m *MockSensor) SetSensorName(sensorName string) {
	m.Sensor.Name = sensorName
}

func (m *MockSensor) SetDeviceId(deviceId string) {
	m.deviceId = deviceId
}

func (m *MockSensor) SetSensorMetaData(sensorMetaData sensor.Sensor) {
	m.Sensor = sensorMetaData
}
