package model

import "time"

type Task struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	Item   string    `json:"item,omitempty"`
	Done   bool      `json:"done,omitempty"`
	DoneAt time.Time `json:"done_at,omitempty"`
}

func NewTask(id string, createdAt int64, updatedAt int64, item string, done bool, doneAt int64) *Task {
	return &Task{
		Id:        id,
		CreatedAt: time.Unix(createdAt, 0).Local(),
		UpdatedAt: time.Unix(updatedAt, 0).UTC(),
		Item:      item,
		Done:      done,
		DoneAt:    time.Unix(doneAt, 0),
	}
}
