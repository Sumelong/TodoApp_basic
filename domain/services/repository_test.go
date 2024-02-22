package services

import (
	"TodoApp_basic/domain/entity"
	"TodoApp_basic/tests"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func Test_Create_Ok(t *testing.T) {

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
	assert.Equal(t, int64(1), res)

}
