package golang_loging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Warn("This is warn")
	logger.Info("This is info")
	logger.Error("This is Error")

}
func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Warn("This is warn")
	logger.Info("This is info")
	logger.Error("This is Error")

}