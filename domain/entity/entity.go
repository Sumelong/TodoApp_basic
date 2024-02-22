package entity

import (
	"TodoApp_basic/services"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id        uuid.UUID `json:"id,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	UpdatedAt int64     `json:"updated_at,omitempty"`

	Item   string `json:"item,omitempty"`
	Done   bool   `json:"done,omitempty"`
	DoneAt int64  `json:"done_at,omitempty"`
}

func NewTask(item string, done bool) *Task {
	var doneAt int64 = 00000
	if done {
		doneAt = time.Now().Unix()
	}
	return &Task{

		Id:        uuid.New(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Item:      services.IsString(item),
		Done:      done,
		DoneAt:    doneAt,
	}

}

func (t *Task) GetID() uuid.UUID {
	return t.Id
}
