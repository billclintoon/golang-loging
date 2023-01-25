package golang_loging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Helllo World")
	logger.Warn("Helllo World")
	logger.Error("Helllo World")
}
