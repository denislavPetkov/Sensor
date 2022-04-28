package sensorfactory

import (
	"fmt"

	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorimpl"
	log "github.com/sirupsen/logrus"
)

type SensorFactory struct{}

type SensorFactoryInterface interface {
	NewSensor(sensorName string, sensorMetaData sensor.Sensor, deviceId string) (sensor.SensorInterface, error)
}

func NewSensorFactory() SensorFactoryInterface {
	return &SensorFactory{}
}

const (
	cpuTemperatureSensorName       = "cpuTempCelsius"
	cpuCoresCountSensorName        = "cpuCoresCount"
	cpuFrequencySensorName         = "cpuFrequency"
	cpuUsagePercentageSensorName   = "cpuUsagePercent"
	gpuMemorySensorName            = "gpuMemoryTotal"
	memoryAvailableSensorName      = "memoryAvailableBytes"
	memoryTotalSensorName          = "memoryTotal"
	memoryUsedSensorName           = "memoryUsedBytes"
	memoryUsedPercentageSensorName = "memoryUsedPercent"
)

var sensorNames = []string{cpuTemperatureSensorName, cpuCoresCountSensorName, cpuFrequencySensorName, cpuUsagePercentageSensorName, gpuMemorySensorName, memoryAvailableSensorName, memoryTotalSensorName, memoryUsedPercentageSensorName, memoryUsedSensorName}

func (f *SensorFactory) NewSensor(sensorName string, sensorMetaData sensor.Sensor, deviceId string) (sensor.SensorInterface, error) {
	switch sensorName {
	case cpuTemperatureSensorName:
		return sensorimpl.NewCpuTemperatureSensor(sensorMetaData, deviceId), nil
	case cpuCoresCountSensorName:
		return sensorimpl.NewCpuCoresCountSensor(sensorMetaData, deviceId), nil
	case cpuFrequencySensorName:
		return sensorimpl.NewCpuFrequencySensor(sensorMetaData, deviceId), nil
	case cpuUsagePercentageSensorName:
		return sensorimpl.NewCpuUsagePercentageSensor(sensorMetaData, deviceId), nil
	case gpuMemorySensorName:
		return sensorimpl.NewGpuMemorySensor(sensorMetaData, deviceId), nil
	case memoryAvailableSensorName:
		return sensorimpl.NewMemoryAvailableSensor(sensorMetaData, deviceId), nil
	case memoryTotalSensorName:
		return sensorimpl.NewMemoryTotalSensor(sensorMetaData, deviceId), nil
	case memoryUsedSensorName:
		return sensorimpl.NewMemoryUsedSensor(sensorMetaData, deviceId), nil
	case memoryUsedPercentageSensorName:
		return sensorimpl.NewMemoryUsedPercentageSensor(sensorMetaData, deviceId), nil
	default:
		log.Error("Invalid sensor name provided")
		return nil, fmt.Errorf("'%v' is not a valid sensor", sensorName)
	}
}

func GetAllSensorNames() []string {
	return sensorNames
}
