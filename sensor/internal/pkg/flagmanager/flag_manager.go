package flagmanager

import (
	"flag"
	"fmt"
	"os"

	"github.com/denislavPetkov/sensor/internal/pkg/flagmanager/customflags"
	"github.com/denislavPetkov/sensor/internal/pkg/model/sensor/sensorfactory"
	"github.com/denislavPetkov/sensor/internal/pkg/printer"
	log "github.com/sirupsen/logrus"
)

var sensorGroupFlags customflags.StringSliceFlag
var helpFlag = flag.Bool("help", false, "lists the commands and their usage")
var deltaDuration = flag.Int(deltaDurationFlagName, 5, "specifies the duration between two sensor measurements")
var totalDuration = flag.Int(totalDurationFlagName, 15, "specifies the total duration after which the program is terminated")
var formatFlag = flag.String("format", "JSON", "specifies the output format (JSON or YAML)")
var webUrlFlag = flag.String("web_hook_url", "", "specifies an url to be called")
var availableSensorsFlag = flag.Bool("get_available_sensors", false, "lists the names of all available senors")

const (
	totalDurationFlagName = "total_duration"
	deltaDurationFlagName = "delta_duration"
)

func Parse() error {
	flag.Var(&sensorGroupFlags, "sensor_group", "specifies the sensor group ('CPU_TEMP', 'CPU_USAGE' and 'MEMORY_USAGE' are default ones)")
	flag.Parse()
	err := validateFlags()
	if err != nil {
		return err
	}
	checkAvailableSensorsFlag()
	printHelp()
	return nil
}

func checkAvailableSensorsFlag() {

	if !*availableSensorsFlag {
		return
	}
	sensorNames := sensorfactory.GetAllSensorNames()

	printer.Println(sensorNames)
	os.Exit(0)
}
func printHelp() {
	currentFlags := GetFlags()
	sensorGroupsNotProvided := currentFlags.SensorGroups == nil
	if *helpFlag || sensorGroupsNotProvided {
		printer.Println("At least 1 sensor group needs to be provided!")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

type Flags struct {
	SensorGroups  []string
	DeltaDuration int
	TotalDuration int
	Format        string
	WebUrl        string
}

func GetFlags() Flags {
	s := Flags{
		SensorGroups:  sensorGroupFlags,
		DeltaDuration: *deltaDuration,
		TotalDuration: *totalDuration + 1,
		Format:        *formatFlag,
		WebUrl:        *webUrlFlag,
	}
	return s
}

func validateFlags() error {
	err := totalDurationFlagValidation(totalDuration)
	if err != nil {
		return err
	}
	err = deltaDurationFlagValidation(deltaDuration)
	if err != nil {
		return err
	}
	return nil
}

func totalDurationFlagValidation(flagValue *int) error {
	return intFlagValidation(flagValue, totalDurationFlagName)
}
func deltaDurationFlagValidation(flagValue *int) error {
	return intFlagValidation(flagValue, deltaDurationFlagName)
}

func intFlagValidation(flagValue *int, flagName string) error {
	if *flagValue <= 0 {
		log.Errorf("Wrong value provided for flag '%v'", flagName)
		return fmt.Errorf("Value for '%v' flag must be a positive number bigger than 0", flagName)
	}
	return nil
}
