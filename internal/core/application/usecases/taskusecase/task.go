package taskusecase

import (
	"TodoApp_basic/internal/core/application/model"
	"TodoApp_basic/internal/core/application/ports/logger"
	"TodoApp_basic/internal/core/application/ports/mappers"
	"TodoApp_basic/internal/core/domain/ports/Idprovider"
	"TodoApp_basic/internal/core/domain/ports/repository"
)

type TaskUseCase struct {
	repo   repository.Repository
	mapper mappers.Mapper
	logger logger.Logger
}

func NewTaskUseCase(
	repo repository.Repository, mapper mappers.Mapper, logger logger.Logger) *TaskUseCase {
	return &TaskUseCase{
		repo:   repo,
		mapper: mapper,
		logger: logger,
	}
}

func (t TaskUseCase) Add(id Idprovider.Id) (string, error) {

	//eTask := mapping.ModelToNewEntity(mTask)
	eTask := t.mapper.ModelToNewEntity(id)
	res, err := t.repo.Create(eTask)
	if err != nil {
		t.logger.Error("task creation error")
		return "", err
	}
	t.logger.Info("task created")
	return res, nil
}

func (t TaskUseCase) Update(mTask *model.Task) (string, error) {

	eTask := t.mapper.ModelToUpdateEntity()
	id, err := t.repo.Update(eTask)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (t TaskUseCase) FindAll() ([]model.Task, error) {

	var mTasks []model.Task
	res, err := t.repo.FindAll()

	for _, en := range res {
		mTasks = append(mTasks, *t.mapper.EntityToModel(&en))
	}

	if err != nil {
		return nil, err
	}
	return mTasks, nil
}

func (t TaskUseCase) FindOne() (*model.Task, error) {

	var mTask *model.Task
	eTask := t.mapper.ModelToAllEntity()

	eTask, err := t.repo.FindOne(eTask)
	if err != nil {
		return nil, err
	}
	mTask = t.mapper.EntityToModel(eTask)
	return mTask, nil
}

func (t TaskUseCase) Remove() (string, error) {

	//map model to entity

	eTask := t.mapper.ModelToAllEntity()

	res, err := t.repo.Remove(eTask)
	if err != nil {
		return "", err
	}

	return res, nil
}
