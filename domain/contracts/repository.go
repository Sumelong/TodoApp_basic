package contracts

import (
	"TodoApp_basic/domain/entity"
)

type Repository interface {
	Create(entity *entity.Base) (int64, error)
	FindAll() ([]entity.Task, error)
	FindBy(entity entity.Base) (entity.Task, error)
	Update(entity entity.Base, Where entity.Base) (int64, error)
	Remove(entity entity.Base) (int64, error)
}
