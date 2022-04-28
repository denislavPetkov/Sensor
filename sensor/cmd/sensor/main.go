package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/denislavPetkov/sensor/internal/pkg/flagmanager"
	"github.com/denislavPetkov/sensor/internal/pkg/logger"
	"github.com/denislavPetkov/sensor/internal/pkg/model/device"
	"github.com/denislavPetkov/sensor/internal/pkg/model/device/deviceregistry"
	"github.com/denislavPetkov/sensor/internal/pkg/model/measurement"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorfactory"
	_ "github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorimpl"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorregistry"
	"github.com/denislavPetkov/sensor/internal/pkg/printer"
	log "github.com/sirupsen/logrus"
)

const modelPath string = "./model.yaml"

func main() {

	err := initialize()
	if err != nil {
		printer.Println(err)
		os.Exit(1)
	}
	err = start()
	if err != nil {
		printer.Println(err)
		os.Exit(1)
	}

}

func start() error {
	log.Info("Program started.")

	err := flagmanager.Parse()
	if err != nil {
		return err
	}

	model, err := ioutil.ReadFile(modelPath)
	if err != nil {
		log.Error(err)
		return fmt.Errorf("Failed to open model file")
	}

	deviceRegistry, err := deviceregistry.NewDeviceRegistry(model)
	if err != nil {
		return err
	}

	devices := deviceRegistry.GetDevices()

	err = registerSensorImplObjects(devices)
	if err != nil {
		return err
	}

	flags := flagmanager.GetFlags()

	err = checkForRegisteredSensors(devices, flags.SensorGroups)
	if err != nil {
		return err
	}

	err = processMeasurements(flags, devices)
	return err
}

func checkForRegisteredSensors(devices []device.Device, sensorGroups []string) error {

	var devicesWithRegisteredSensorsCount int = 0

	for _, device := range devices {
		if device.DoRegisteredSensorsExist(sensorGroups) {
			devicesWithRegisteredSensorsCount++
		}
	}

	if devicesWithRegisteredSensorsCount == 0 {
		return fmt.Errorf("No sensors available from neither device")
	}

	return nil
}

func registerSensorImplObjects(devices []device.Device) error {
	sensorFactory := sensorfactory.NewSensorFactory()

	for i := 0; i < len(devices); i++ {
		deviceId := devices[i].GetId()
		sensorRegistry := sensorregistry.NewSensorRegistry()

		var sensorsToRegister []sensor.SensorInterface
		sensorsInDevice := devices[i].GetSensors()

		for _, sensor := range sensorsInDevice {
			sensorImpl, err := sensorFactory.NewSensor(sensor.Name, sensor, deviceId)
			if err != nil {
				return err
			}
			sensorsToRegister = append(sensorsToRegister, sensorImpl)
		}

		sensorRegistry.RegisterSensors(sensorsToRegister)
		devices[i].SetRegistry(sensorRegistry)

	}
	return nil
}

func processMeasurements(flags flagmanager.Flags, devices []device.Device) error {
	doneChannel := time.After(time.Duration(flags.TotalDuration) * time.Second)
	deltaChannel := time.Tick(time.Duration(flags.DeltaDuration) * time.Second)

	for {
		select {
		case <-deltaChannel:
			err := getAndPrintMeasurements(flags, devices)
			if err != nil {
				return err
			}
		case <-doneChannel:
			log.Info("Program exited with no errors")
			return nil
		}
	}
}

func getAndPrintMeasurements(flags flagmanager.Flags, devices []device.Device) error {

	for _, device := range devices {
		measurements, err := device.GetMeasurements(flags.SensorGroups)
		if err != nil {
			return err
		}
		err = printMeasurements(measurements, flags)
		if err != nil {
			return err
		}
	}
	return nil
}

func printMeasurements(measurements []measurement.Measurement, flags flagmanager.Flags) error {
	for _, measurement := range measurements {
		err := measurement.PrintMeasurement(flags.Format, flags.WebUrl)
		if err != nil {
			return err
		}
	}
	return nil
}

func initialize() error {
	var f os.File
	err := logger.SetUpLogger(&f)
	if err != nil {
		return fmt.Errorf("Failed to create log file: '%v'", err)
	}

	log.Info("Application initialized successfully")
	return nil
}
