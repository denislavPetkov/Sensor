package sensorfactory

import (
	"testing"

	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSensorfactory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sensorfactory Suite")
}

const (
	invalidSensorName                    = "Invalid"
	deviceId                             = "1"
	sensorName                           = "cpuTempCelsius"
	sensorId                             = "1"
	sensorGroup       sensor.SensorGroup = "CPU_TEMP"
)

var (
	sensorMetaData = sensor.Sensor{
		Name:         sensorName,
		Id:           sensorId,
		SensorGroups: []sensor.SensorGroup{sensorGroup},
	}
)
