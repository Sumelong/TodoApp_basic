package logging

import (
	"TodoApp_basic/adapters/logger"
	"log/slog"
	"os"
)

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLog(logFile *os.File) *SlogLogger {

	jsonHandler := slog.NewJSONHandler(logFile, nil)
	slogLog := slog.New(jsonHandler)
	slogLog.Info("slog logger Initiated")

	return &SlogLogger{logger: slogLog}
}

func (l *SlogLogger) Info(format string, args ...interface{}) {
	l.logger.Info(format, args...)
}

func (l *SlogLogger) Warn(format string, args ...interface{}) {
	l.logger.Warn(format, args...)
}

func (l *SlogLogger) Error(format string, args ...interface{}) {
	l.logger.Error(format, args...)
}

func (l *SlogLogger) Debug(format string, args ...interface{}) {
	l.logger.Debug(format, args)
}

func (l *SlogLogger) WithFields(fields logger.Fields) logger.Logger {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}
	log := l.logger.With(f...)
	return &SlogLogger{logger: log}
}

func (l *SlogLogger) WithError(err error) logger.Logger {
	var log = l.logger.With(err.Error())
	return &SlogLogger{logger: log}
}
