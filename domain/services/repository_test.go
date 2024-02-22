package services

import (
	"TodoApp_basic/domain/entity"
	"TodoApp_basic/tests"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestRepository_Create(t *testing.T) {

	//setup
	dsn, db, err := tests.TestInit()
	if err != nil {
		t.Error(err)
	}
	defer tests.TestCleanUp(dsn, db)

	//arrange
	task := entity.NewTask("cook", false)
	expected := task.GetID()
	repo := NewRepository(db)

	//act
	res, err := repo.Create(task)

	//assert
	assert.NoError(t, err)
	assert.Equal(t, expected, res)

}

func TestRepository_FindAll(t *testing.T) {
	//setup
	dsn, db, err := tests.TestInit()
	if err != nil {
		t.Error(err)
	}
	defer tests.TestCleanUp(dsn, db)

	//arrange
	repo := NewRepository(db)
	var tasks []entity.Task
	for i := 0; i < 10; i++ {
		task := entity.NewTask(fmt.Sprintf("task-%d", i), false)
		_, err = repo.Create(task)
		if err != nil {
			t.Error(err)
		}
		tasks = append(tasks, *task)
	}

	//act
	res, err := repo.FindAll()

	//assert

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.ObjectsAreEqual(tasks, res)
	assert.Equal(t, tasks[0], res[0])
	assert.NotEqual(t, tasks[0], res[1])

}

func TestRepository_FindBy(t *testing.T) {
	//setup
	dsn, db, err := tests.TestInit()
	if err != nil {
		t.Error(err)
	}
	defer tests.TestCleanUp(dsn, db)

	//arrange
	repo := NewRepository(db)
	var tasks []entity.Task
	for i := 0; i < 10; i++ {
		task := entity.NewTask(fmt.Sprintf("task-%d", i), false)
		_, err = repo.Create(task)
		if err != nil {
			t.Error(err)
		}
		tasks = append(tasks, *task)
	}

	//act
	Where := entity.Task{Id: tasks[0].Id}
	res, err := repo.FindBy(&Where)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, tasks[0].Item, res.Item)

}

func TestRepository_Update(t *testing.T) {
	//setup
	dsn, db, err := tests.TestInit()
	if err != nil {
		t.Error(err)
	}
	defer tests.TestCleanUp(dsn, db)

	//arrange
	repo := NewRepository(db)
	var tasks []entity.Task
	for i := 0; i < 10; i++ {
		task := entity.NewTask(fmt.Sprintf("task-%d", i), false)
		_, err = repo.Create(task)
		if err != nil {
			t.Error(err)
		}
		tasks = append(tasks, *task)
	}

	//act
	//Where := entity.Task{Id: tasks[0].Id}
	newTask := entity.UpdateTask(tasks[0].Id, "newTask", true)
	res, err := repo.Update(newTask)

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, newTask.Id, res)

	updated, _ := repo.FindBy(newTask)
	assert.NotEqual(t, tasks[0], &updated)
	assert.True(t, updated.Done == true)
	assert.Equal(t, updated.Item, newTask.Item)

}

func TestRepository_Remove(t *testing.T) {

	//setup
	dsn, db, err := tests.TestInit()
	if err != nil {
		t.Error(err)
	}
	defer tests.TestCleanUp(dsn, db)

	//arrange
	repo := NewRepository(db)
	var tasks []entity.Task
	for i := 0; i < 10; i++ {
		task := entity.NewTask(fmt.Sprintf("task-%d", i), false)
		_, err = repo.Create(task)
		if err != nil {
			t.Error(err)
		}
		tasks = append(tasks, *task)
	}

	//act
	res, err := repo.Remove(&tasks[0])

	//assert
	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	updated, _ := repo.FindBy(&tasks[0])
	assert.Empty(t, updated)

}
