package tests

import (
	"TodoApp_basic/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockDb struct {
	mock.Mock
	tasks map[string]entity.Task
}

func (d *MockDb) Exec(eTask entity.Task) (int64, error) {
	d.tasks[eTask.Id] = eTask
	args := d.Called()
	return int64(args.Int(0)), args.Error(1)
}
func (d *MockDb) Query(args ...any) (entity.Task, error) {
	args = append(args, args...)
	arg := d.Called(args)
	return arg.Get(0).(entity.Task), arg.Error(1)
}
