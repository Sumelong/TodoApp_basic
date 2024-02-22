package task

import (
	"TodoApp_basic/application/mapper"
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/contracts"
	"TodoApp_basic/domain/entity"
	"database/sql"
)

type Task struct {
	db   *sql.DB
	repo contracts.Repository
}

func NewTask(db *sql.DB, repo contracts.Repository) *Task {
	return &Task{
		db:   db,
		repo: repo,
	}
}

func (t Task) Add(task *model.Task) (string, error) {

	entity := mapper.ToEntity(task)
	id, err := t.repo.Create(entity)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t Task) Update(task *model.Task) (string, error) {

	ent := mapper.ToEntity(task)
	id, err := t.repo.Update(ent)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t Task) FindAll() ([]entity.Task, error) {

	res, err := t.repo.FindAll()

	if err != nil {
		return nil, err
	}
	return res, nil
}
