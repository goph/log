package logrus

import "github.com/sirupsen/logrus"

// Logger is a wrapper around Logrus.
type Logger struct {
	Logger logrus.FieldLogger
}

// NewLogger returns a new Logger.
func NewLogger() *Logger {
	return &Logger{logrus.New()}
}

// Debug logs on Debug level.
func (l *Logger) Debug(args ...interface{}) {
	msg, ctx := parseLog(args...)

	lo := l.Logger

	if len(ctx) > 0 {
		lo = lo.WithFields(ctx)
	}

	lo.Debug(msg)
}

// Info logs on Info level.
func (l *Logger) Info(args ...interface{}) {
	msg, ctx := parseLog(args...)

	lo := l.Logger

	if len(ctx) > 0 {
		lo = lo.WithFields(ctx)
	}

	lo.Info(msg)
}

// Error logs on Error level.
func (l *Logger) Error(args ...interface{}) {
	msg, ctx := parseLog(args...)

	lo := l.Logger

	if len(ctx) > 0 {
		lo = lo.WithFields(ctx)
	}

	lo.Error(msg)
}

func parseLog(args ...interface{}) (string, map[string]interface{}) {
	var msg string
	var ctx map[string]interface{}

	if len(args) <= 2 {
		for _, value := range args {
			if value, ok := value.(string); ok {
				msg = value
				continue
			}

			if value, ok := value.(map[string]interface{}); ok {
				ctx = value
				continue
			}
		}
	}

	return msg, ctx
}
