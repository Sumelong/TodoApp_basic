package entity

import (
	"TodoApp_basic/internal/core/application/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTask(t *testing.T) {

	//arrange
	id := services.NewId()

	//act
	eTask := NewTask("task", false, id)

	//assert
	assert.NotEmpty(t, eTask.Id)
	assert.Equal(t, eTask.Id, id.GetId())

}
