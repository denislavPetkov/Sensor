package sensorregistry

import (
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
)

var _ = Describe("Sensorregistry", func() {

	var testSensorMetaData sensor.Sensor = sensor.Sensor{
		Id:           sensorId,
		Name:         sensorName,
		SensorGroups: []sensor.SensorGroup{sensorGroup},
	}

	var testSensor *MockSensor = &MockSensor{
		Sensor:   testSensorMetaData,
		deviceId: deviceId,
	}

	sensorRegistry := NewSensorRegistry()

	Describe("RegisterSensors()", func() {
		Context("When RegisterSensors(sensor), then the provided sensors are registered in the repository", func() {
			It("Should register the sensors in repository", func() {
				sensorRegistry.RegisterSensors([]sensor.SensorInterface{testSensor})
				expected := sensorRegistry.GetSensorByName(sensorName)
				Expect(expected).To(Equal(testSensor))
			})
		})
	})

	Describe("GetSensorsByGroup()", func() {
		Context("When GetSensorsByGroup(), then a slice with all registered sensors from the specific is returned", func() {
			It("Should return all registered sensors from wanted sensor group", func() {
				expected := sensorRegistry.GetSensorsByGroup(sensorGroup)
				Expect(expected[0]).To(Equal(testSensor))
			})
		})
	})

})
