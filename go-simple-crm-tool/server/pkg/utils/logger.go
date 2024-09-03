package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger(level logrus.Level, format string, output string) {
	Logger = logrus.New()
	Logger.SetLevel(level)

	switch format {
	case "json":
		Logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	switch output {
	case "stdout":
		Logger.SetOutput(os.Stdout)
	case "stderr":
		Logger.SetOutput(os.Stderr)
	default:
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			Logger.SetOutput(os.Stdout)
			Logger.Warn("Failed to log to file, using default stdout")
		} else {
			Logger.SetOutput(file)
		}
	}
}

func Info(fields logrus.Fields, message string) {
	Logger.WithFields(fields).Info(message)
}

func Warn(fields logrus.Fields, message string) {
	Logger.WithFields(fields).Warn(message)
}

func Fatal(fields logrus.Fields, message string) {
	Logger.WithFields(fields).Fatal(message)
}
