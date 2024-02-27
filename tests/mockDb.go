package tests

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockDb struct {
	mock.Mock
	open sql.DB
}

func (d *MockDb) Open() *sql.DB {
	return d.Called().Get(0).(*sql.DB)
}
