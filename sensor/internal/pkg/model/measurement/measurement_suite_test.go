package measurement

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMeasurement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Measurement Suite")
}

const (
	value      = "10"
	sensorId   = "1"
	deviceId   = "1"
	jsonFormat = "JSON"
	yamlFormat = "YAML"
)
