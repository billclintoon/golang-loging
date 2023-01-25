package golang_loging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("name", "bill").Info("Hello World")
	logger.WithField("username", "bill").
		Logger.WithField("username", "bill clinton").
		Info("Hello bill")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "bill",
		"name":     "bill clinton",
	}).Info("hello bill")
}
