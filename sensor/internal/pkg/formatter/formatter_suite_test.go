package formatter

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFormat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format Suite")
}

const (
	invalidFormat             = "XML"
	deviceId                  = "1"
	temp              float64 = 50
	cpuTempSensorName         = "cpuTempCelsius"
	cpuTempSensorId           = "1"
)
