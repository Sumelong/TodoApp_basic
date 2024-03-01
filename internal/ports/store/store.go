package store

import (
	"database/sql"
)

type DB interface {
	Begin() (*sql.Tx, error)
}
