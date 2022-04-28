package sensorimpl

import (
	"testing"

	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSensorimpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sensorimpl Suite")
}

var (
	testSensor = sensor.Sensor{
		Id: testSensorId,
	}

	expectedSensorInterface sensor.SensorInterface
)

const testSensorId = "1"
const testDeviceId = "1"
