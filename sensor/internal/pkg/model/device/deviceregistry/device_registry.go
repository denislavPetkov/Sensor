package deviceregistry

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/denislavPetkov/sensor/internal/pkg/model/device"
	"gopkg.in/yaml.v2"
)

var once sync.Once

type DeviceRegistry struct {
	Devices []device.Device `yaml:"devices"`
}

func NewDeviceRegistry(model []byte) (DeviceRegistryInterface, error) {
	var registry *DeviceRegistry
	var err error

	once.Do(func() {
		registry = &DeviceRegistry{}
		err = yaml.Unmarshal(model, registry)

	})

	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("Failed to create device registry")
	}

	return registry, nil
}

func (m *DeviceRegistry) GetDevices() []device.Device {
	return m.Devices
}

type DeviceRegistryInterface interface {
	GetDevices() []device.Device
}
