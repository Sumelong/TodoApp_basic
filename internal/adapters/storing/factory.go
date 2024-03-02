package storing

import (
	"TodoApp_basic/internal/adapters/storing/sqlite"
	"TodoApp_basic/internal/core/application/ports/query"
	"errors"
)

// storing Instances
const (
	InstanceSqlite int = iota
	InstancePostgres
)

var ErrInvalidQueryInstance = errors.New("invalid storing instance")

func NewQuery(instance int) (*query.Query, error) {

	switch instance {
	case InstanceSqlite:
		return sqlite.Query(), nil
	default:
		return nil, ErrInvalidQueryInstance

	}
}
