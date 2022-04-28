package sensorfactory

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sensorfactory", func() {

	sensorFactory := NewSensorFactory()

	Describe("NewSensor()", func() {
		Context("When NewSensor(), then a new sensor is returned", func() {
			It("Should return correct sensor", func() {
				actual, _ := sensorFactory.NewSensor(sensorName, sensorMetaData, deviceId)
				Expect(actual.GetName()).To(Equal(sensorName))
				Expect(actual.GetSensorGroups()[0]).To(Equal(sensorGroup))
			})
		})
		Context("When NewSensor() with incorrect sensor name, then an error is returned", func() {
			It("Should return an error", func() {
				actual, err := sensorFactory.NewSensor(invalidSensorName, sensorMetaData, deviceId)
				Expect(actual).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

	})

})
