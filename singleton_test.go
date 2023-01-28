package belajar_golang_loging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello")
	logrus.Error("Hello Err")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Hello Info")
}
