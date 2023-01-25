package golang_loging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")
	logrus.Warn("Hello world")
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World")
	logrus.Warn("Hello World")
}
