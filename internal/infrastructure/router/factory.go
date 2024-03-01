package router

import (
	"TodoApp_basic/adapters/logger"
	"database/sql"
	"time"
)

type Server interface {
	Run()
}

type Port int64

const (
	InstanceCustomServer int = iota
	InstanceServerMux
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
		return NewRouter(log, port, StoreSql, ctxTimeout), nil
	case InstanceServerMux:
		return NewMuxServer(log, port, StoreSql, ctxTimeout), nil

	//case InstanceGin:
	//	return newGinServer(log, dbNoSQL, validator, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
