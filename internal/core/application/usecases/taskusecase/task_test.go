package taskusecase

import (
	"TodoApp_basic/internal/adapters/storing"
	"TodoApp_basic/internal/core/application/model"
	"TodoApp_basic/internal/core/application/ports/logger"
	"TodoApp_basic/internal/core/application/services"
	"TodoApp_basic/internal/core/application/services/mapping"
	"TodoApp_basic/internal/core/application/usecases/taskusecase/mockers"
	"TodoApp_basic/internal/core/domain/entity"
	"TodoApp_basic/internal/core/domain/ports/Idprovider"
	"TodoApp_basic/internal/core/domain/repository"
	"TodoApp_basic/internal/infrastructure/logging"
	"TodoApp_basic/tests"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InitTasks(t *testing.T, db *sql.DB) ([]model.Task, []entity.Task, logger.Logger, error) {

	var (
		id  Idprovider.Id
		err error
	)

	log, err := logging.NewLoggerFactory(logging.InstanceSlogLogger)
	if err != nil {
		t.Error(err)
	}

	qry, err := storing.NewQuery(storing.InstanceSqlite)

	if err != nil {
		t.Error(t)
	}

	mockRepo := repository.NewTaskRepository(db, *qry)

	var mTasks []model.Task
	var eTasks []entity.Task
	for i := 0; i < 10; i++ {
		newTask := model.NewTask(fmt.Sprintf("cook-%d", i), false)
		id = services.NewId()
		m := mapping.NewMap(newTask)
		uc := NewTaskUseCase(mockRepo, m, log)
		mTasks = append(mTasks, *newTask)
		_, err = uc.Add(id)
	}

	for _, task := range mTasks {
		m := mapping.NewMap(&task)
		eTasks = append(eTasks, *m.ModelToAllEntity())
	}

	return mTasks, eTasks, log, err
}

func TestTaskService_Add(t *testing.T) {

	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	log, err := logging.NewLoggerFactory(logging.InstanceSlogLogger)
	if err != nil {
		t.Error(err)
	}

	//arrange
	newTask := model.Task{Item: "cook", Done: false} //NewTask("cook", false)
	mockRepo := new(mockers.MockRepository)

	id := services.NewId()
	m := mapping.NewMap(&newTask)

	//act
	uc := NewTaskUseCase(mockRepo, m, log)
	eTask := m.ModelToNewEntity(id)

	mockRepo.On("Create", eTask).Return(id.GetId(), nil)

	res, err := uc.Add(id)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, id.GetId(), res)

	mockRepo.AssertExpectations(t)

}

func TestTaskService_Update_Item(t *testing.T) {
	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mTasks, _, log, err := InitTasks(t, db)
	m := mapping.NewMap(&mTasks[1])
	mockRepo := new(mockers.MockRepository)
	uc := NewTaskUseCase(mockRepo, m, log)
	uTask := model.Task{Id: mTasks[1].Id, Item: "Walk", Done: false}
	eTask := m.ModelToUpdateEntity()
	mockRepo.On("Update", eTask).Return(mTasks[1].Id, nil)

	//act

	res, err := uc.Update(&uTask)
	log.Info("test log okay")

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, uTask.Id, res)
	assert.NotEqual(t, uTask.Item, mTasks[1].Item) // not correctly tested

	mockRepo.AssertExpectations(t)

}

func TestTaskService_Update_Done(t *testing.T) {
	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mTasks, _, log, err := InitTasks(t, db)
	m := mapping.NewMap(&mTasks[1])
	mockRepo := new(mockers.MockRepository)
	uc := NewTaskUseCase(mockRepo, m, log)
	uTask := model.Task{Id: mTasks[1].Id, Item: "Walk", Done: true}
	eTask := m.ModelToUpdateEntity()
	mockRepo.On("Update", eTask).Return(mTasks[1].Id, nil)

	//act

	res, err := uc.Update(&uTask)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, uTask.Id, res)
	assert.NotEqual(t, uTask.Done, mTasks[1].Done) // not correctly tested

	mockRepo.AssertExpectations(t)

}

func TestTaskService_FindAll(t *testing.T) {
	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mTasks, eTask, log, err := InitTasks(t, db)
	m := mapping.NewMap(&mTasks[1])
	mockRepo := new(mockers.MockRepository)
	uc := NewTaskUseCase(mockRepo, m, log)

	mockRepo.On("FindAll").Return(eTask, nil)

	//act

	res, err := uc.FindAll()

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(mTasks), len(res), fmt.Sprintf("expected %d, got %d", len(mTasks), len(res)))

	mockRepo.AssertExpectations(t)

}

func TestTaskService_FindOne(t *testing.T) {
	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mTasks, eTasks, log, err := InitTasks(t, db)
	m := mapping.NewMap(&mTasks[1])
	mockRepo := new(mockers.MockRepository)
	uc := NewTaskUseCase(mockRepo, m, log)

	expected := &eTasks[1]

	mockRepo.On("FindOne", expected).Return(expected, nil)

	//act

	res, err := uc.FindOne()

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, expected.Id, res.Id)

	mockRepo.AssertExpectations(t)

}

func TestTaskService_Removed(t *testing.T) {

	//setup
	dns, db, err := tests.TestInit()
	if err != nil {
		t.Fatal(err)
	}
	defer tests.TestCleanUp(dns, db)

	//arrange
	mTasks, eTasks, log, err := InitTasks(t, db)
	m := mapping.NewMap(&mTasks[1])
	mockRepo := new(mockers.MockRepository)
	uc := NewTaskUseCase(mockRepo, m, log)

	expected := &eTasks[1]

	mockRepo.On("Remove", expected).Return(expected.Id, nil)

	//act

	res, err := uc.Remove()

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, expected.Id, res)

	mockRepo.AssertExpectations(t)

}
