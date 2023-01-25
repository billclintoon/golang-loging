package golang_loging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("sample hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Warn("This is warn")
	logger.Info("This is info")
	logger.Error("This is Error")
	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
}
