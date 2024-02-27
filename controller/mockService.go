package controller

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/contracts"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
	repo contracts.Repository
}

func (c *MockTaskService) Add(mTask *model.Task) (string, error) {
	args := c.Called(mTask)
	return args.String(0), args.Error(1)
}

func (c *MockTaskService) Update(mTask *model.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockTaskService) FindAll() ([]model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockTaskService) FindOne(mTask *model.Task) (*model.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (c *MockTaskService) Remove(mTask *model.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewMockTaskService(repo contracts.Repository) *MockTaskService {
	return &MockTaskService{
		repo: repo,
	}
}
