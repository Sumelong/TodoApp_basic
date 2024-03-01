package model

import (
	"TodoApp_basic/application/usecase"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Item   string    `json:"item,omitempty"`
	Done   bool      `json:"done,omitempty"`
	DoneAt time.Time `json:"done_at,omitempty"`
}

func NewTask(item string, done bool) *Task {
	var doneAt time.Time
	if done {
		doneAt = time.Now()
	}
	return &Task{
		Id:        uuid.NewString(),
		CreatedAt: time.Now(),
		Item:      usecase.IsString(item),
		Done:      done,
		DoneAt:    doneAt,
	}

}

func UpdateTask(id string, item string, done bool) *Task {
	var doneAt time.Time
	if done {
		doneAt = time.Now()
	}
	return &Task{
		Id:        id,
		UpdatedAt: time.Now(),
		Item:      usecase.IsString(item),
		Done:      done,
		DoneAt:    doneAt,
	}

}

func (t *Task) GetID() string {
	return t.Id
}
