package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryUsedSensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &memoryUsedSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
	})

	Describe("getUsedMemory()", func() {
		testSensor := &memoryUsedSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
		Context("When getUsedMemory(), then correct used memory in bytes returned", func() {
			It("Should return correct used memory in bytes", func() {
				actual, err := testSensor.getUsedMemory()
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
