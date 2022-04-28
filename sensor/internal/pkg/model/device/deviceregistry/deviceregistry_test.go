package deviceregistry

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("DeviceRegistry", func() {

	err := yaml.Unmarshal([]byte(testModel), testDeviceRegistry)
	Expect(err).NotTo(HaveOccurred())

	Describe("CreateDeviceRegistry()", func() {
		Context("When CreateDeviceRegistry(), then a new device registry is created and returned", func() {
			expectedDeviceRegistry := testDeviceRegistry
			It("Should return create and return a device registry", func() {
				actual, err := NewDeviceRegistry([]byte(testModel))
				Expect(err).NotTo(HaveOccurred())
				Expect(actual).To(Equal(expectedDeviceRegistry))
			})
		})
	})

	Describe("GetDevices()", func() {
		Context("When GetDevices(), then correct devices are returned", func() {
			It("Should return correct slice with devices", func() {
				actual := testDeviceRegistry.GetDevices()
				Expect(actual[0].Name).To(Equal(deviceName))
				Expect(actual[0].Id).To(Equal(deviceId))

			})
		})
	})

})
