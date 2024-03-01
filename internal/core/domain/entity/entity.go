package entity

import (
	"TodoApp_basic/application/usecase"
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

func NewTask(item string, done bool) *Task {
	var doneAt int64 = 00000
	if done {
		doneAt = time.Now().Unix()
	}
	return &Task{

		Id:        uuid.New().String(),
		CreatedAt: time.Now().Unix(),
		Item:      usecase.IsString(item),
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
		Item:      usecase.IsString(item),
		Done:      done,
		DoneAt:    doneAt,
	}

}

func (t *Task) GetID() string {
	return t.Id
}
