package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	fileName  = "sensor.log"
	fileFlags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	fileMode  = 0644
)

func SetUpLogger(f *os.File) error {
	defer f.Close()

	var err error
	f, err = os.OpenFile(fileName, fileFlags, fileMode)

	if err != nil {
		return err
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(f)

	return nil
}
