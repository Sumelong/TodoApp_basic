package logging

import (
	"TodoApp_basic/adapters/logger"
	"log"
	"os"
)

type StdLogger struct {
	logger   *log.Logger
	logLevel string
}

func NewStdLog(logFile *os.File) *StdLogger {

	var myLogger StdLogger
	level := myLogger.logLevel
	myLogger.logger = log.New(logFile, level, log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	return &myLogger
}

func (l *StdLogger) Info(format string, args ...interface{}) {
	l.logLevel = "[info]"
	l.logger.Printf(format, args...)
}
func (l *StdLogger) Warn(format string, args ...interface{}) {
	l.logLevel = "[warn]"
	l.logger.Printf(format, args...)
}
func (l *StdLogger) Error(format string, args ...interface{}) {
	l.logLevel = "[Error]"
	l.logger.Printf(format, args...)
}

func (l *StdLogger) Debug(format string, args ...interface{}) {
	l.logLevel = "[Debug]"
	l.logger.Printf(format, args...)
}

func (l *StdLogger) WithFields(fields logger.Fields) {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}
	l.logger.Println(f...)
}

func (l *StdLogger) WithError(err error) {
	l.logger.Println(err.Error())
}
