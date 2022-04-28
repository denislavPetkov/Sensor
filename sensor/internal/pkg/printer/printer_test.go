package printer

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Printer", func() {

	Describe("CallWebHookUrl()", func() {
		var server *ghttp.Server
		var url string
		var expectedMeasurement interface{}

		BeforeEach(func() {

			timeStamp := time.Now()
			expectedMeasurement = struct {
				MeasuredAt time.Time `yaml:"measuredAt" json:"measuredAt"`
				Value      string    `yanl:"value" json:"value"`
				SensorId   string    `yaml:"sensorId" json:"sensorId"`
				DeviceId   string    `yaml:"deviceId" json:"deviceId"`
			}{
				MeasuredAt: timeStamp,
				Value:      sensorValue,
				SensorId:   sensorId,
				DeviceId:   deviceId,
			}

			server = ghttp.NewServer()
			url = server.URL() + "/measurement"
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/measurement"),
					ghttp.VerifyJSONRepresenting(expectedMeasurement),
					ghttp.RespondWithJSONEncoded(http.StatusOK, nil),
				),
			)
		})

		AfterEach(func() {
			server.Close()
		})

		Context("When called", func() {
			It("Should send a measurement to the url", func() {
				jsonMeasurement, err := json.Marshal(expectedMeasurement)
				Expect(err).NotTo(HaveOccurred())

				CallWebHookUrl(string(jsonMeasurement), url)
			})
		})
	})

})
