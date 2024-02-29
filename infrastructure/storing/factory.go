package storing

import (
	"TodoApp_basic/adapters/logger"
	"database/sql"
	"errors"
)

const (
	InstanceSqlite = iota
)

var (
	errInvalidStoreInstance = errors.New("invalid store instance")
)

func NewStoreFactory(instance int, dns string, logger logger.Logger) (*sql.DB, error) {

	switch instance {
	case InstanceSqlite:
		return NewSqlite(dns, logger)

	default:
		return nil, errInvalidStoreInstance
	}
}
