package log

/*
LevelLogger attaches a level to the log.

Supported method signatures:

	logger.Level("Message")
	logger.Level("Message", map[string]interface{}{"key": "value"})
	logger.Level(map[string]interface{}{"key": "value"})
*/
type LevelLogger interface {
	// Debug logs on Debug level.
	Debug(args ...interface{})

	// Info logs on Info level.
	Info(args ...interface{})

	// Error logs on Error level.
	Error(args ...interface{})
}
