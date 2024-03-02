package mappers

import (
	"TodoApp_basic/internal/core/application/model"
	"TodoApp_basic/internal/core/domain/entity"
	"TodoApp_basic/internal/core/domain/ports/Idprovider"
)

type Mapper interface {
	ModelToNewEntity(id Idprovider.Id) *entity.Task
	ModelToUpdateEntity() *entity.Task
	ModelToAllEntity() *entity.Task
	EntityToModel(entity *entity.Task) *model.Task
}
