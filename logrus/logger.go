package logrus

import (
	"github.com/goph/log"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logger logrus.FieldLogger
}

func (l *Logger) WithField(key string, value interface{}) log.StructuredLogger {
	return &Logger{l.Logger.WithField(key, value)}
}

func (l *Logger) WithFields(fields map[string]interface{}) log.StructuredLogger {
	return &Logger{l.Logger.WithFields(logrus.Fields(fields))}
}

func (l *Logger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}
