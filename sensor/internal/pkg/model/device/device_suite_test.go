package device

import (
	"testing"

	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDevice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Device Suite")
}

const (
	deviceId          = "1"
	deviceName        = "1"
	cpuTempSensorName = "cpuTempSensor"
)

var (
	testSensor = sensor.Sensor{
		Name: cpuTempSensorName,
	}

	mockRegistry = &mockSensorRegistry{}
)

type mockSensorRegistry struct{}

func (s *mockSensorRegistry) RegisterSensors(sensors []sensor.SensorInterface) {

}

func (s *mockSensorRegistry) GetSensorByName(sensorName string) sensor.SensorInterface {
	return nil
}

func (s *mockSensorRegistry) GetSensorsByGroup(desiredSensorGroup string) []sensor.SensorInterface {
	return nil
}
