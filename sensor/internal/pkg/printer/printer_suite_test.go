package printer

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrinter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Printer Suite")
}

const (
	sensorValue = "99"
	sensorId    = "1"
	deviceId    = "1"
)
