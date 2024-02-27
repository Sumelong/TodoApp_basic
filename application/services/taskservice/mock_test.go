package taskservice

import (
	"TodoApp_basic/domain/entity"
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
	db *sql.DB
}

func (r *MockRepository) Create(entity *entity.Task) (string, error) {
	args := r.Called(entity)
	return args.Get(0).(string), args.Error(1)
}

func (r *MockRepository) FindAll() ([]entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MockRepository) FindBy(entity *entity.Task) (*entity.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *MockRepository) Update(entity *entity.Task) (string, error) {
	args := r.Called(entity)
	return args.Get(0).(string), args.Error(1)
}

func (r *MockRepository) Remove(entity *entity.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}
