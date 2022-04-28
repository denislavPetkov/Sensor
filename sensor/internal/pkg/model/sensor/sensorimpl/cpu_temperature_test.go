package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CpuTemperatureSensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &cpuTemperatureSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
	})

	Describe("getTemperature()", func() {
		testSensor := &cpuTemperatureSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
		Context("When getTemperature(), then correct cpu temperature is returned", func() {
			It("Should return correct cpu temperature", func() {
				actual, err := testSensor.getTemperature()
				Expect(actual).NotTo(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetName()", func() {
		Context("When GetName(), then new correct sensor name is returned", func() {
			It("Should return correct sensor name", func() {
				actual := expectedSensorInterface.GetName()
				Expect(actual).To(Equal(testSensor.Name))
			})
		})
	})

	Describe("GetMeasurement()", func() {
		Context("When GetMeasurement(), then correct measurement is returned", func() {
			It("Should return correct measurement", func() {
				actual, _ := expectedSensorInterface.GetMeasurement()
				Expect(actual.Value).NotTo(Equal(""))
				Expect(actual.MeasuredAt).NotTo(Equal(""))
				Expect(actual.SensorId).To(Equal(testSensorId))
				Expect(actual.DeviceId).To(Equal(testDeviceId))
			})
		})
	})

})
