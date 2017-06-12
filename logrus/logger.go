package logrus

import (
	"github.com/goph/log"
	"github.com/sirupsen/logrus"
)

// Logger is a wrapper around Logrus.
type Logger struct {
	Logger logrus.FieldLogger
}

// WithField allows to add a key-value pair to the logger as context.
//
// It returns a new instance in order to avoid modifing the parent context.
func (l *Logger) WithField(key string, value interface{}) log.StructuredLogger {
	return &Logger{l.Logger.WithField(key, value)}
}

// WithFields allows to add a set of key-value pairs to the logger as context.
//
// It returns a new instance in order to avoid modifing the parent context.
func (l *Logger) WithFields(fields map[string]interface{}) log.StructuredLogger {
	return &Logger{l.Logger.WithFields(logrus.Fields(fields))}
}

// Debug logs on Debug level.
func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

// Info logs on Info level.
func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

// Error logs on Error level.
func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}
