package sensorimpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CpuCoresCountSensor", func() {

	BeforeEach(func() {
		expectedSensorInterface = &cpuCoresCountSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
	})

	Describe("NewCpuCoresCountSensor()", func() {
		Context("When NewCpuCoresCountSensor(), then new correct sensor is returned", func() {
			It("Should return correct sensor", func() {
				actual := NewCpuCoresCountSensor(testSensor, testDeviceId)
				Expect(actual).To(Equal(expectedSensorInterface))
			})
		})
	})

	Describe("getCoresCount()", func() {
		testSensor := &cpuCoresCountSensor{
			sensorMetadata: testSensor,
			deviceId:       testDeviceId,
		}
		Context("When getCoresCount(), then correct number of cores is returned", func() {
			It("Should return correct number of cores", func() {
				actual, err := testSensor.getCoresCount()
				Expect(actual).NotTo(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetName()", func() {
		Context("When GetName(), then correct sensor name is returned", func() {
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
