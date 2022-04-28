package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GpuMemorySensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &gpuMemorySensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
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
