package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryAvailableSensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &memoryAvailableSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
	})

	Describe("getAvailableMemory()", func() {
		testSensor := &memoryAvailableSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
		Context("When getAvailableMemory(), then correct available is returned", func() {
			It("Should return correct available memory in bytes", func() {
				actual, err := testSensor.getAvailableMemory()
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
