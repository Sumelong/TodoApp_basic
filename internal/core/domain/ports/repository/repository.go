package repository

import (
	"TodoApp_basic/internal/core/domain/entity"
)

type Repository interface {
	Create(eTask *entity.Task) (string, error)
	Update(eTask *entity.Task) (string, error)
	FindAll() ([]entity.Task, error)
	FindOne(eTask *entity.Task) (*entity.Task, error)
	Remove(eTask *entity.Task) (string, error)
}
