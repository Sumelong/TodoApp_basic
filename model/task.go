package model

import "time"

type Task struct {
	Base
	Item   string    `json:"item"`
	Done   bool      `json:"done"`
	DoneAt time.Time `json:"done_at"`
}
