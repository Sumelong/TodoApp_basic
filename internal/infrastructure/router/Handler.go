package router

import "C"
import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/controller"
	"encoding/json"
	"net/http"
)

var (
	mTask model.Task
	item  string
)

type TaskHandler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func SetHandlers() *Router {
	route := NewRouter()
	c := controller.NewTaskController()

	route.GET("/", c.FindAll)
	route.GET("/", c.FindAllTaskHandler)
	route.POST("/add", c.CreateTaskHandler)

	route.PATCH("/PAT", c.CreateTaskHandler)
	//route.POST("/delete/{id}", c.FindAllTaskHandler)

	return route
}
