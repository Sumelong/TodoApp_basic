package taskservice

import (
	"TodoApp_basic/application/mapper"
	"TodoApp_basic/application/model"
	"TodoApp_basic/tests"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskService_Add(t *testing.T) {

	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(MockRepository)
	taskService := NewTaskService(db, mockRepo)
	newTask := model.NewTask("cook", false)

	e := mapper.ModelToEntity(newTask)
	mockRepo.On("Create", e).Return(newTask.Id, nil)

	//act
	res, err := taskService.Add(newTask)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, newTask.Id, res)

	mockRepo.AssertExpectations(t)

}

func TestTaskService_Update(t *testing.T) {
	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(MockRepository)
	taskService := NewTaskService(db, mockRepo)

	newTask := model.NewTask("cook", false)
	//addedId, _ := taskService.Add(newTask)

	e := mapper.ModelToEntity(newTask)
	mockRepo.On("Update", e).Return(newTask.Id, nil)

	//act
	updateTask := model.UpdateTask(newTask.Id, "Walk", false)
	res, err := taskService.Update(updateTask)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, newTask.Id, res)

	mockRepo.AssertExpectations(t)

}
