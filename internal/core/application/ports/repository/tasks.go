package repository

import (
	"TodoApp_basic/internal/core/application/model"
)

type Task interface {
	Add(mTask *model.Task) (string, error)
	Update(mTask *model.Task) (string, error)
	FindAll() ([]model.Task, error)
	FindOne(mTask *model.Task) (*model.Task, error)
	Remove(mTask *model.Task) (string, error)
}
