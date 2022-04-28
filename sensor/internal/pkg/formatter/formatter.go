package formatter

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Format string

const (
	Json Format = "JSON"
	Yaml Format = "YAML"
)

func FormatMeasurement(measurement interface{}, selectedFormat string) (string, error) {

	switch selectedFormat {
	case string(Json):
		return toJson(measurement)
	case string(Yaml):
		return toYaml(measurement)
	default:
		log.Warn("Wrong format has been choosen!")
		return "", fmt.Errorf("Invalid format selected!")
	}

}

func toYaml(measurement interface{}) (string, error) {
	formattedData, err := yaml.Marshal(measurement)
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("A problem has occurred during the coversion to YAML format")
	}
	formattedMeasurement := string(formattedData)

	return formattedMeasurement, nil
}

func toJson(measurement interface{}) (string, error) {
	formattedData, err := json.Marshal(measurement)
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("A problem has occurred during the coversion to JSON format")
	}
	formattedMeasurement := string(formattedData)

	return formattedMeasurement, nil
}
