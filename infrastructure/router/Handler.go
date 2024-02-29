package router

import (
	"TodoApp_basic/controller"
	"encoding/json"
	"net/http"
)

type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statusCode, data := h(r)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Init() *Router {
	r := NewRouter()
	c := controller.NewTaskController()

	r.GET("/task", c.FindAllTaskHandler)

	route.HandleFunc("/", controllers.Show)
	route.HandleFunc("/add", controllers.Add).Methods("POST")
	route.HandleFunc("/delete/{id}", controllers.Delete)
	route.HandleFunc("/complete/{id}", controllers.Complete)

	return route
}
