package taskservice

import (
	"TodoApp_basic/app_tests"
	"TodoApp_basic/application/mapper"
	"TodoApp_basic/application/model"
	"TodoApp_basic/domain/entity"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTaskService_Add(t *testing.T) {

	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	taskService := NewTaskService(mockRepo)
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

func TestTaskService_Update_Item(t *testing.T) {
	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	taskService := NewTaskService(mockRepo)

	newTask := model.NewTask("cook", false)
	//_, _ = taskService.Add(newTask)

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

func TestTaskService_Update_Done(t *testing.T) {
	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	taskService := NewTaskService(mockRepo)

	mTask := model.NewTask("cook", false)
	//_, _ = taskService.Add(newTask)

	eTask := mapper.ModelToEntity(mTask)
	mockRepo.On("Update", eTask).Return(mTask.Id, nil)

	//act
	updateTask := model.UpdateTask(mTask.Id, "cook", false)
	res, err := taskService.Update(updateTask)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, mTask.Id, res)

	mockRepo.AssertExpectations(t)
}

func TestTaskService_FindAll(t *testing.T) {
	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	taskService := NewTaskService(mockRepo)

	var mTasks []model.Task
	for i := 0; i < 10; i++ {
		newTask := model.NewTask(fmt.Sprintf("cook-%d", i), false)
		mTasks = append(mTasks, *newTask)
	}

	var eTasks []entity.Task
	for _, mTask := range mTasks {
		eTask := mapper.ModelToEntity(&mTask)
		eTasks = append(eTasks, *eTask)
	}

	mockRepo.On("FindAll").Return(eTasks, nil)

	//act
	res, err := taskService.FindAll()

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	//assert.Equal(t, mTasks, res)
	assert.Equal(t, mTasks[1].Item, res[1].Item)

	mockRepo.AssertExpectations(t)

}

func TestTaskService_FindOne(t *testing.T) {
	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	//mockRepo :=  repository.NewTaskRepository(db) //new(MockRepository)

	taskService := NewTaskService(mockRepo)

	var mTasks []model.Task
	for i := 0; i < 10; i++ {
		newTask := model.NewTask(fmt.Sprintf("cook-%d", i), false)
		mTasks = append(mTasks, *newTask)
	}

	//where := &model.Task{Id: mTasks[1].Id}

	eTask := mapper.ModelToEntity(&mTasks[1])
	mockRepo.On("FindOne", eTask).Return(eTask, nil)

	//act
	res, err := taskService.FindOne(&mTasks[1])

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	//assert.Equal(t, mTasks[1], res)
	assert.Equal(t, mTasks[1].Item, res.Item)

	mockRepo.AssertExpectations(t)
}

func TestTaskService_Removed(t *testing.T) {

	//setup
	dns, db, err := app_tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer app_tests.TestCleanUp(dns, db)

	//arrange
	mockRepo := new(app_tests.MockRepository)
	//mockRepo :=  repository.NewTaskRepository(db) //new(MockRepository)

	taskService := NewTaskService(mockRepo)

	var mTasks []model.Task
	for i := 0; i < 10; i++ {
		newTask := model.NewTask(fmt.Sprintf("cook-%d", i), false)
		mTasks = append(mTasks, *newTask)
	}

	//where := &model.Task{Id: mTasks[1].Id}

	eTask := mapper.ModelToEntity(&mTasks[1])
	mockRepo.On("Remove", eTask).Return(eTask.Id, nil)

	//act
	res, err := taskService.Remove(&mTasks[1])

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, mTasks[1].Id, res)

	mockRepo.AssertExpectations(t)
}
