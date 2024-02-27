package mapper

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
	"time"
)

func ModelToEntity(model *model.Task) *entity.Task {
	var doneAt time.Time
	if model.Done {
		doneAt = time.Now()
	}

	return &entity.Task{
		Id:        model.Id,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: model.UpdatedAt.Unix(),
		Item:      model.Item,
		Done:      model.Done,
		DoneAt:    doneAt.Unix(),
	}
}

func EntityToModel(entity *entity.Task) *model.Task {
	return &model.Task{
		Id:        entity.Id,
		CreatedAt: time.Unix(entity.CreatedAt, 0),
		UpdatedAt: time.Unix(entity.UpdatedAt, 0),
		Item:      entity.Item,
		Done:      entity.Done,
		DoneAt:    time.Time{},
	}
}
