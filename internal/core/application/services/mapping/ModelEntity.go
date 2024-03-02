package mapping

import (
	"TodoApp_basic/internal/core/application/model"
	"TodoApp_basic/internal/core/domain/entity"
	"TodoApp_basic/internal/core/domain/ports/Idprovider"
	"time"
)

type MapToEntity struct {
	task *model.Task
}

func NewMap(task *model.Task) *MapToEntity {
	return &MapToEntity{task: task}
}

func (m MapToEntity) ModelToNewEntity(id Idprovider.Id) *entity.Task {

	return entity.NewTask(m.task.Item, m.task.Done, id)
}

func (m MapToEntity) ModelToUpdateEntity() *entity.Task {
	return entity.UpdateTask(m.task.Id, m.task.Item, m.task.Done)
}

func (m MapToEntity) ModelToAllEntity() *entity.Task {

	return &entity.Task{
		Id:        m.task.Id,
		CreatedAt: m.task.CreatedAt.Unix(),
		UpdatedAt: m.task.UpdatedAt.Unix(),
		Item:      m.task.Item,
		Done:      m.task.Done,
		DoneAt:    m.task.DoneAt.Unix(),
	}
}

func (m MapToEntity) EntityToModel(entity *entity.Task) *model.Task {
	return &model.Task{
		Id:        entity.Id,
		CreatedAt: time.Unix(entity.CreatedAt, 0),
		UpdatedAt: time.Unix(entity.UpdatedAt, 0),
		Item:      entity.Item,
		Done:      entity.Done,
		DoneAt:    time.Unix(entity.DoneAt, 0),
	}
}
