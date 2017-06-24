package logrus

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	levelField = "level"
	msgField   = "msg"

	debugLevel = "debug"
	infoLevel  = "info"
	errorLevel = "error"
)

// Logger is a wrapper around Logrus.
type Logger struct {
	Logger logrus.FieldLogger
}

// NewLogger returns a new Logger.
func NewLogger() *Logger {
	return &Logger{logrus.New()}
}

// Log implements the logger interface.
func (l *Logger) Log(keyvals ...interface{}) error {
	kvlen := len(keyvals)

	if kvlen%2 != 0 {
		keyvals = append(keyvals, "(MISSING)")
		kvlen++
	}

	// TODO: default level?
	level := "info"
	var msg string
	ctx := map[string]interface{}{}

	// Note: unsafe string assertion
	for i := 0; i < kvlen-1; i += 2 {
		switch keyvals[i] {
		case msgField:
			msg = keyvals[i+1].(string)

		case levelField:
			if s, ok := keyvals[i+1].(fmt.Stringer); ok {
				level = s.String()

			} else {
				level = keyvals[i+1].(string)
			}

		default:
			ctx[keyvals[i].(string)] = keyvals[i+1]
		}
	}

	lo := l.Logger.WithFields(ctx)

	switch level {
	case debugLevel:
		lo.Debug(msg)

	default:
	case infoLevel:
		lo.Info(msg)

	case errorLevel:
		lo.Error(msg)
	}

	return nil
}
