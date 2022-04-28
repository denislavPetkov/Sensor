package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MemoryTotalSensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &memoryTotalSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
	})

	Describe("getTotalMemory()", func() {
		testSensor := &memoryTotalSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
		Context("When getTotalMemory(), then correct total memory is returned", func() {
			It("Should return correct total memory in gigabytes", func() {
				actual, err := testSensor.getTotalMemory()
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
