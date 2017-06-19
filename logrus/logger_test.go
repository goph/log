package logrus_test

import (
	"testing"

	"github.com/goph/log/logrus"
	_logrus "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

func TestLogger_Debug_Empty(t *testing.T) {
	testEmpty("Debug", t)
}

func TestLogger_Info_Empty(t *testing.T) {
	testEmpty("Info", t)
}

func TestLogger_Error_Empty(t *testing.T) {
	testEmpty("Error", t)
}

func TestLogger_Debug_Message(t *testing.T) {
	testMessage("Debug", t)
}

func TestLogger_Info_Message(t *testing.T) {
	testMessage("Info", t)
}

func TestLogger_Error_Message(t *testing.T) {
	testMessage("Error", t)
}

func TestLogger_Debug_Context(t *testing.T) {
	testContext("Debug", t)
}

func TestLogger_Info_Context(t *testing.T) {
	testContext("Info", t)
}

func TestLogger_Error_Context(t *testing.T) {
	testContext("Error", t)
}

func TestLogger_Debug_MessageAndContext(t *testing.T) {
	testMessageAndContext("Debug", t)
}

func TestLogger_Info_MessageAndContext(t *testing.T) {
	testMessageAndContext("Info", t)
}

func TestLogger_Error_MessageAndContext(t *testing.T) {
	testMessageAndContext("Error", t)
}

func TestLogger_Debug_Invalid(t *testing.T) {
	testInvalid("Debug", t)
}

func TestLogger_Info_Invalid(t *testing.T) {
	testInvalid("Info", t)
}

func TestLogger_Error_Invalid(t *testing.T) {
	testInvalid("Error", t)
}

func testEmpty(level string, t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	invoke(logger, level, t)

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "" {
		t.Error("message is expected to be empty")
	}

	if len(lastEntry.Data) != 0 {
		t.Error("context is expected to be empty")
	}
}

func testMessage(level string, t *testing.T, args ...interface{}) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	invoke(logger, level, t, "Test message")

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "Test message" {
		t.Errorf(`message is expected to be "Test message", received: %s`, lastEntry.Message)
	}

	if len(lastEntry.Data) != 0 {
		t.Error("context is expected to be empty")
	}
}

func testContext(level string, t *testing.T, args ...interface{}) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	invoke(logger, level, t, map[string]interface{}{"key": "value"})

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

func testMessageAndContext(level string, t *testing.T, args ...interface{}) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	invoke(logger, level, t, "Test message", map[string]interface{}{"key": "value"})

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

func testInvalid(level string, t *testing.T) {
	testLogger, hook := test.NewNullLogger()
	testLogger.Level = _logrus.DebugLevel
	logger := &logrus.Logger{
		Logger: testLogger,
	}

	invoke(logger, level, t, "invalid 1", "invalid 2", "invalid 3")

	lastEntry := hook.LastEntry()

	if lastEntry.Message != "" {
		t.Error("message is expected to be empty")
	}

	if len(lastEntry.Data) != 0 {
		t.Error("context is expected to be empty")
	}
}

func invoke(logger *logrus.Logger, level string, t *testing.T, args ...interface{}) {
	switch level {
	case "Debug":
		logger.Debug(args...)

	case "Info":
		logger.Info(args...)

	case "Error":
		logger.Error(args...)

	default:
		t.Fatalf("unknown logging level: %s", level)
	}
}
