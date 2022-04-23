package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	return logger
}
