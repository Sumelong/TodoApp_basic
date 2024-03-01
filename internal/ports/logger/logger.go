package logger

type Logger interface {
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Debug(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
	WithError(err error) Logger
}

type Fields map[string]interface{}
