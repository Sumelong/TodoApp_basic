package controller

import (
	"TodoApp_basic/adapters/usecase"
	"TodoApp_basic/application/model"
	"encoding/json"
	"html/template"
	"net/http"
)

type TaskController struct {
	uc usecase.Task
}

var (
	view  = template.Must(template.ParseFiles("./views/index.html"))
	mTask = model.Task{}
)

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func NewTaskController(service usecase.Task) *TaskController {
	return &TaskController{uc: service}
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
	res, err := c.uc.Add(&mTask)
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

func (c *TaskController) FindAllTaskHandler(w http.ResponseWriter, r *http.Request) {

	// get task from data store
	res, err := c.uc.FindAll()
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

func (c *TaskController) FindAll(r *http.Request) (statusCode int, data map[string]interface{}) {

	return 0, nil
}

func (c *TaskController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := c.uc.Add(r.Body)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
