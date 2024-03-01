package logging

import (
	"TodoApp_basic/adapters/logger"
	"errors"
	"log"
)

const (
	InstanceStdLogger int = iota
	InstanceSlogLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {
	loggingFile, err := openLogFile("./logs.log")
	if err != nil {
		log.Fatal(err)
	}

	switch instance {
	case InstanceStdLogger:
		return NewSlogLog(loggingFile), nil
	case InstanceSlogLogger:
		return NewSlogLog(loggingFile), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
