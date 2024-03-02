package entity

import (
	"TodoApp_basic/internal/core/domain/ports/Idprovider"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        string
	CreatedAt int64
	UpdatedAt int64

	Item   string
	Done   bool
	DoneAt int64
}

func NewTask(item string, done bool, id Idprovider.Id) *Task {
	var doneAt int64

	if done {
		doneAt = time.Now().Unix()
	}
	//set id to get
	//id.SetId()

	return &Task{

		Id:        id.GetId(), //uuid.NewString(),
		CreatedAt: time.Now().Unix(),
		Item:      item,
		Done:      done,
		DoneAt:    doneAt,
	}
}

func UpdateTask(id string, item string, done bool) *Task {
	var doneAt int64 = 00000
	if done {
		doneAt = time.Now().Unix()
	}
	return &Task{
		Id:        id,
		UpdatedAt: time.Now().Unix(),
		Item:      item,
		Done:      done,
		DoneAt:    doneAt,
	}

}

func (t *Task) GetID() string {
	return t.Id
}

func (t *Task) setID() {
	t.Id = uuid.NewString()
}
