package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)
var log *logrus.Logger

func LogInit(logLevel string) {
	log = logrus.New()

	level := logrus.DebugLevel
	log.SetLevel(logrus.ErrorLevel)
	switch {
	case logLevel == "debug":
		level = logrus.DebugLevel
		case logLevel == "info":
			level = logrus.InfoLevel
			case logLevel == "error":
				level = logrus.ErrorLevel
				default:
					level = logrus.DebugLevel
	}
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}

func Info(v ...interface{}) {
	log.Info(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Debug(v ...interface{}) {
	log.Debug(v)
}


