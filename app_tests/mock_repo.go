package app_tests

import (
	"TodoApp_basic/domain/entity"
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
	db *sql.DB
}

func (r *MockRepository) Create(eTask *entity.Task) (string, error) {
	args := r.Called(eTask)
	return args.Get(0).(string), args.Error(1)
}

func (r *MockRepository) FindAll() ([]entity.Task, error) {
	args := r.Called()
	return args.Get(0).([]entity.Task), args.Error(1)
}

func (r *MockRepository) FindOne(eTask *entity.Task) (*entity.Task, error) {
	args := r.Called(eTask)
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (r *MockRepository) Update(eTask *entity.Task) (string, error) {
	args := r.Called(eTask)
	return args.Get(0).(string), args.Error(1)
}

func (r *MockRepository) Remove(eTask *entity.Task) (string, error) {
	args := r.Called(eTask)
	return args.String(0), args.Error(1)
}
