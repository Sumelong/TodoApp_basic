package contracts

import (
	"TodoApp_basic/domain/entity"
)

type Repository interface {
	Create(eTask *entity.Task) (string, error)
	FindAll() ([]entity.Task, error)

	//FindOne get the entity to find by id and return the entity found or error
	FindOne(eTask *entity.Task) (*entity.Task, error)
	Update(eTask *entity.Task) (string, error)
	Remove(eTask *entity.Task) (string, error)
}
