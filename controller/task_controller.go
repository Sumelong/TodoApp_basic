package controller

import (
	"TodoApp_basic/application/contracts"
	"TodoApp_basic/application/model"
	"encoding/json"
	"net/http"
)

type TaskController struct {
	Service contracts.Task
}

func NewTaskController(service contracts.Task) *TaskController {
	return &TaskController{Service: service}
}

// CreateTaskHandler handles HTTP requests to create a new user
func (c *TaskController) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var mTask model.Task

	err := json.NewDecoder(r.Body).Decode(&mTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Validate user data (for example, check required fields)
	if mTask.Item == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Save user to the database
	res, err := c.Service.Add(&mTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
