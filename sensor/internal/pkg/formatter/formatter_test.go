package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"gopkg.in/yaml.v2"
)

var _ = Describe("Format", func() {

	var testMeasurement interface{}

	Describe("GetFormattedString() functionality", func() {

		timeStamp := time.Now()
		testMeasurement = struct {
			MeasuredAt time.Time
			Value      string
			SensorId   string
			DeviceId   string
		}{
			MeasuredAt: timeStamp,
			Value:      fmt.Sprintf("%v", temp),
			SensorId:   cpuTempSensorId,
			DeviceId:   deviceId,
		}

		expected := func(selectedFormat Format) string {

			switch selectedFormat {
			case Json:
				formatedData, err := json.Marshal(testMeasurement)
				Expect(err).NotTo(HaveOccurred())
				return string(formatedData)
			case Yaml:
				formatedData, err := yaml.Marshal(testMeasurement)
				Expect(err).NotTo(HaveOccurred())
				return string(formatedData)
			default:
				return ""
			}
		}

		DescribeTable("different formats", func(measurement interface{}, selectedFormat string, expected string) {
			actual, _ := FormatMeasurement(measurement, selectedFormat)
			Expect(actual).To(Equal(expected))
		},

			Entry("when format is JSON", testMeasurement, string(Json), expected(Json)),
			Entry("when format is YAML", testMeasurement, string(Yaml), expected(Yaml)),
			Entry("when format is invalid", testMeasurement, invalidFormat, expected(invalidFormat)),
		)
	})

})
