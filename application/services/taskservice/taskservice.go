package taskservice

import (
	"TodoApp_basic/application/mapper"
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/contracts"
)

type TaskService struct {
	repo contracts.Repository
}

func NewTaskService(repo contracts.Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t TaskService) Add(mTask *model.Task) (string, error) {

	eTask := mapper.ModelToNewEntity(mTask)
	id, err := t.repo.Create(eTask)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TaskService) Update(mTask *model.Task) (string, error) {

	eTask := mapper.ModelToUpdateEntity(mTask)
	id, err := t.repo.Update(eTask)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TaskService) FindAll() ([]model.Task, error) {

	var mTasks []model.Task
	res, err := t.repo.FindAll()

	for _, en := range res {
		mTasks = append(mTasks, *mapper.EntityToModel(&en))
	}

	if err != nil {
		return nil, err
	}
	return mTasks, nil
}

func (t TaskService) FindOne(mTask *model.Task) (*model.Task, error) {

	//var mTask *model.Task
	eTask := mapper.ModelToEntity(mTask)

	eTask, err := t.repo.FindOne(eTask)
	if err != nil {
		return nil, err
	}
	mTask = mapper.EntityToModel(eTask)
	return mTask, nil
}

func (t TaskService) Remove(mTask *model.Task) (string, error) {
	//map model to entity
	eTask := mapper.ModelToEntity(mTask)

	res, err := t.repo.Remove(eTask)
	if err != nil {
		return "", err
	}
	mTask = mapper.EntityToModel(eTask)
	return res, nil
}
