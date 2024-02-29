package repository

import (
	"TodoApp_basic/app_tests"
	"TodoApp_basic/domain/entity"
	"fmt"
	"testing"
)

func BenchmarkTaskRepository_Create(b *testing.B) {
	dsn, db, err := app_tests.TestInit()
	if err != nil {
		b.Error(err)
	}
	defer app_tests.TestCleanUp(dsn, db)

	//arrange
	task := entity.NewTask("cook", false)
	repo := NewTaskRepository(db)

	for i := 0; i < b.N; i++ {
		_, err = repo.Create(task)
		if err != nil {
			b.Error(err)
		}

	}
}

func BenchmarkTaskRepository_FindAll(b *testing.B) {
	//setup
	dsn, db, err := app_tests.TestInit()
	if err != nil {
		b.Error(err)
	}
	defer app_tests.TestCleanUp(dsn, db)

	//arrange
	repo := NewTaskRepository(db)
	var tasks []entity.Task
	for i := 0; i < 10; i++ {
		task := entity.NewTask(fmt.Sprintf("taskservice-%d", i), false)
		_, err = repo.Create(task)
		if err != nil {
			b.Error(err)
		}
		tasks = append(tasks, *task)
	}

	//act

	for i := 0; i < b.N; i++ {
		_, err = repo.FindAll()
		if err != nil {
			b.Error(err)
		}
	}
}
