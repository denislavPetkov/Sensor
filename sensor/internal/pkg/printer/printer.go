package printer

import (
	"bytes"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Println(data ...interface{}) {
	fmt.Println(data...)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func CallWebHookUrl(measurement string, webUrl string) {

	jsonMeasurement := []byte(measurement)

	resp, err := http.Post(webUrl, "application/json", bytes.NewBuffer(jsonMeasurement))
	if err != nil {
		log.Error(err)
		Printf("Failed to send measurement to url - '%v'\n", webUrl)
		return
	}

	log.Infoln("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	defer resp.Body.Close()

}
