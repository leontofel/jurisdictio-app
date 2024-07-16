package pkg

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.DebugLevel)
}

func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}
