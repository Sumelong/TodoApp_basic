package contracts

import (
	"TodoApp_basic/domain/entity"
)

type Repository interface {
	Create(entity *entity.Task) (string, error)
	FindAll() ([]entity.Task, error)

	//FindBy get the entity to find by id and return the entity found or error
	FindBy(entity *entity.Task) (*entity.Task, error)
	Update(entity *entity.Task) (string, error)
	Remove(entity *entity.Task) (string, error)
}
