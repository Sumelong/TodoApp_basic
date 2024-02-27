package taskservice

import (
	"TodoApp_basic/application/mapper"
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/contracts"
	"TodoApp_basic/domain/entity"
	"database/sql"
)

type TaskService struct {
	db   *sql.DB
	repo contracts.Repository
}

func NewTaskService(db *sql.DB, repo contracts.Repository) *TaskService {
	return &TaskService{
		db:   db,
		repo: repo,
	}
}

func (t TaskService) Add(task *model.Task) (string, error) {

	e := mapper.ModelToEntity(task)
	id, err := t.repo.Create(e)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TaskService) Update(task *model.Task) (string, error) {

	ent := mapper.ModelToEntity(task)
	id, err := t.repo.Update(ent)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TaskService) FindAll() ([]entity.Task, error) {

	res, err := t.repo.FindAll()

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t TaskService) FindOne(task *model.Task) (*model.Task, error) {

	var m *model.Task
	e := mapper.ModelToEntity(task)

	res, err := t.repo.FindBy(e)
	if err != nil {
		return m, err
	}
	m = mapper.ToModel(res)
	return m, nil
}
