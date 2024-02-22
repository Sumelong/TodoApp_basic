package mapper

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
)

func ToEntity(model *model.Task) *entity.Task {
	return entity.NewTask(model.Item, model.Done)
}
