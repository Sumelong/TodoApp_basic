package mapper

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
)

func ToModel(entity *entity.Task) *model.Task {
	return model.NewTask(entity.Id, entity.CreatedAt, entity.UpdatedAt, entity.Item, entity.Done, entity.DoneAt)
}
