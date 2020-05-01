package slog

import "errors"

// A global variable so that log functions can be directly accessed
var log Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	// Debug has verbose message
	Debug = "debug"
	// Info is default log level
	Info = "info"
	// Warn is for logging messages about possible issues
	Warn = "warn"
	// Error is for logging errors
	Error = "error"
	// Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

// Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Infow(message string, args ...interface{})

	Warnw(message string, args ...interface{})

	Errorw(message string, args ...interface{})

	Fatalw(message string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

// GetInstance get logger instance
func GetInstance() Logger {
	return log
}

// NewLogger returns an instance of logger
func NewLogger(config Configuration) {
	log = newZapLogger(config)
}

// Debugf debug format
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof info format
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf warning format
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf error format
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Infow info write
func Infow(message string, args ...interface{}) {
	log.Infow(message, args...)
}

// Warnw warning write
func Warnw(message string, args ...interface{}) {
	log.Warnw(message, args...)
}

// Errorw error write
func Errorw(message string, args ...interface{}) {
	log.Errorw(message, args...)
}

// Fatalw fatal write
func Fatalw(message string, args ...interface{}) {
	log.Fatalw(message, args...)
}

// WithFields add custom fields
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
