package mapper

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
	"time"
)

func ToModel(entity *entity.Task) *model.Task {
	return &model.Task{
		Id:        entity.Id,
		CreatedAt: time.Unix(entity.CreatedAt, 0).Local(),
		UpdatedAt: time.Unix(entity.UpdatedAt, 0).Local(),
		Item:      entity.Item,
		Done:      entity.Done,
		DoneAt:    time.Unix(entity.DoneAt, 01).Local(),
	}
}

func FromModel(model *model.Task) *entity.Task {
	return &entity.Task{
		Id:        model.Id,
		CreatedAt: model.CreatedAt.Unix(),
		UpdatedAt: 0,
		Item:      "",
		Done:      false,
		DoneAt:    0,
	}
}
