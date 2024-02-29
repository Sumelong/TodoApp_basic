package router

import (
	"TodoApp_basic/adapters/logger"
	"database/sql"
	"errors"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceCustomServer int = iota
	InstanceGorillaMux
	InstanceGin
)

func NewWebServerFactory(
	instance int,
	log logger.Logger,
	StoreSql *sql.DB,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceCustomServer:
		return NewRouter(log, port), nil
	//case InstanceGin:
	//	return newGinServer(log, dbNoSQL, validator, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
