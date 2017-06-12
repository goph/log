package log

// Context holds the structured data for a log as a key-value map.
type Context map[string]interface{}

// StructuredLogger allows to attach context to a logger.
type StructuredLogger interface {
	// WithField allows to add a key-value pair to the logger as context.
	//
	// It returns a new instance in order to avoid modifing the parent context.
	WithField(key string, value interface{}) StructuredLogger

	// WithFields allows to add a set of key-value pairs to the logger as context.
	//
	// It returns a new instance in order to avoid modifing the parent context.
	WithFields(fields map[string]interface{}) StructuredLogger
}

// LevelLogger attaches a level to the log.
type LevelLogger interface {
	// Debug logs on Debug level.
	Debug(args ...interface{})

	// Info logs on Info level.
	Info(args ...interface{})

	// Error logs on Error level.
	Error(args ...interface{})
}

// StdLogger declares an interface for the Logger in the standard library.
type StdLogger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
}
