package logrus_test

import (
	"testing"

	"github.com/goph/log/logrus"
	_logrus "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestLogger_Empty(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	logger.Log()

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "" {
		t.Error("message is expected to be empty")
	}

	if len(lastEntry.Data) != 0 {
		t.Error("context is expected to be empty")
	}
}

func TestLogger_Message(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	logger.Log("msg", "Test message")

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "Test message" {
		t.Errorf(`message is expected to be "Test message", received: %s`, lastEntry.Message)
	}

	if len(lastEntry.Data) != 0 {
		t.Error("context is expected to be empty")
	}
}

func TestLogger_Context(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	logger.Log("key", "value")

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "" {
		t.Error("message is expected to be empty")
	}

	if len(lastEntry.Data) == 0 {
		t.Error("context is not expected to be empty")
	}

	if val, ok := lastEntry.Data["key"]; !ok || val != "value" {
		t.Errorf(`context is expected to contain "value" under "key", received: %v`, val)
	}
}

func TestLogger_MessageAndContext(t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	logger.Log("msg", "Test message", "key", "value")

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "Test message" {
		t.Errorf(`message is expected to be "Test message", received: %s`, lastEntry.Message)
	}

	if len(lastEntry.Data) == 0 {
		t.Error("context is not expected to be empty")
	}

	if val, ok := lastEntry.Data["key"]; !ok || val != "value" {
		t.Errorf(`context is expected to contain "value" under "key", received: %v`, val)
	}
}

func TestLogger_Level_Debug(t *testing.T) {
	testLevel(_logrus.DebugLevel, t)
}

func TestLogger_Level_Info(t *testing.T) {
	testLevel(_logrus.InfoLevel, t)
}

func TestLogger_Level_Error(t *testing.T) {
	testLevel(_logrus.ErrorLevel, t)
}

func testLevel(level _logrus.Level, t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	logger.Log("level", level.String())

	lastEntry := hook.LastEntry()

	if lastEntry.Level != level {
		t.Errorf(`message is expected to have %s level, received: %s`, level.String(), lastEntry.Level.String())
	}
}
