package measurement

import (
	"encoding/json"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Measurements", func() {

	var expected Measurement
	var timeStamp time.Time

	BeforeEach(func() {
		timeStamp = time.Now()
		expected = Measurement{timeStamp, value, sensorId, deviceId}
	})

	Describe("CreateMeasurement()", func() {
		Context("When called", func() {
			It("Returns a measurement", func() {
				actual := CreateMeasurement(timeStamp, value, sensorId, deviceId)
				Expect(actual).To(Equal(expected))
			})
		})

	})

	Describe("formatMeasurement()", func() {
		Context("When called with a json format", func() {
			var expectedJson []byte
			BeforeEach(func() {
				var err error
				expectedJson, err = json.Marshal(expected)
				Expect(err).NotTo(HaveOccurred())
			})
			It("Formats the measurement into a json and returns it as a string", func() {
				actual, _ := expected.formatMeasurement(jsonFormat)
				Expect(actual).To(BeEquivalentTo(expectedJson))
			})
		})

		Context("When called with a yaml format", func() {
			var expectedYaml []byte
			BeforeEach(func() {
				var err error
				expectedYaml, err = yaml.Marshal(expected)
				Expect(err).NotTo(HaveOccurred())
			})
			It("Formats the measurement into a yaml and returns it as a string", func() {
				actual, _ := expected.formatMeasurement(yamlFormat)
				Expect(actual).To(BeEquivalentTo(expectedYaml))
			})
		})
	})

	Describe("PrintMeasurements", func() {
		Context("When called without a web url", func() {
			It("Prints the measurements", func() {
				err := expected.PrintMeasurement(jsonFormat, "")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		var server *ghttp.Server
		var url string

		BeforeEach(func() {
			server = ghttp.NewServer()
			url = server.URL() + "/measurement"
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/measurement"),
				),
			)
		})
		AfterEach(func() {
			server.Close()
		})

		Context("When called with a web url in yaml format", func() {
			It("Prints the measurements and sends it to the url in json", func() {
				err := expected.PrintMeasurement(yamlFormat, url)
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

})
