package log

// Logger is a simple logger interface based on go-kit.
//
// The same interface has been proposed to add to the standard library.
// See https://docs.google.com/document/d/1shW9DZJXOeGbG9Mr9Us9MiaPqmlcVatD_D8lrOXRNMU/edit#
type Logger interface {
	Log(keyvals ...interface{}) error
}
