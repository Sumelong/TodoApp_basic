package logging

import (
	"TodoApp_basic/internal/core/application/ports/logger"
	"TodoApp_basic/internal/core/application/services/makedirectory"
	"errors"
	"log"
	"path/filepath"
)

const (
	InstanceStdLogger int = iota
	InstanceSlogLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instance int) (logger.Logger, error) {

	path := filepath.Join("logs")

	err := makedirectory.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	loggingFile, err := openLogFile("logs/logs.log")
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
