package contracts

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
)

type Task interface {
	Add(task *model.Task) (string, error)
	Update(task *model.Task) (string, error)
	FindAll() ([]entity.Task, error)

	FindOne(task *model.Task) (model.Task, error)
}
