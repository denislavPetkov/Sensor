package deviceregistry

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeviceRegistry Suite")
}

var (
	testDeviceRegistry DeviceRegistryInterface = &DeviceRegistry{}
)

const (
	deviceId   = "3"
	deviceName = "device_name"
)

const testModel = `---
devices:
  - id: 3
    name: device_name
    description: my laptop device
    sensors:
      - id: 11
        name: cpuTempCelsius
        description: Measures CPU temp Celsius
        unit: C
        sensorGroups:
          - CPU_TEMP
      - id: 12
        name: cpuUsagePercent
        description: Measures CPU usage percent
        unit: "%"
        sensorGroups:
          - CPU_USAGE
`
