package belajar_golang_loging

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
	fmt.Println("Hello Logger")
}

func TestLogLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello Log 2")
	logger.Error("Hello Log Err 2")
	logger.Warn("Hello Log Warn 2")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Log Info Json")
	logger.Warn("Hello Log Warn Json")
	logger.Error("Hello Log Err Json")
}

func TestLogWithField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("user", "mazeko").Info("Hello")
	//multiple
	logger.WithFields(logrus.Fields{
		"user": "eko",
		"name": "mazeko",
	}).Info("Hello")
}

func TestEntryLog(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("user", "mazeko")
	entry.Info("Hello Entry")
}
