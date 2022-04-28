package device

import (
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Device", func() {

	testDevice := &Device{
		Id:             deviceId,
		Name:           deviceName,
		Sensors:        []sensor.Sensor{testSensor},
		SensorRegistry: mockRegistry,
	}

	Describe("GetSensors()", func() {
		Context("When GetSensors(), then device's sensors are returned", func() {
			It("Should return correct device sensors", func() {
				actual := testDevice.GetSensors()
				Expect(actual[0]).To(Equal(testSensor))
			})
		})
	})

	Describe("GetMeasurement()", func() {
		Context("When GetMeasurement(), then sensor's measurements from specified sensor group are returned", func() {
			sensorGroup := "CPU_TEMP"
			It("Should return correct measurements", func() {
				_, err := testDevice.GetMeasurements([]string{sensorGroup})
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetId()", func() {
		Context("When GetId(), then correct device id is returned", func() {
			It("Should return correct device id", func() {
				actual := testDevice.GetId()
				Expect(actual).To(Equal(deviceId))
			})
		})
	})

})
