package main

import (
	"TodoApp_basic/application/services/taskservice"
	"TodoApp_basic/controller"
	"TodoApp_basic/domain/repository"
	"TodoApp_basic/infrastructure/storing"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
)

func main() {

	_, db, err := storing.DbInit()
	if err != nil {
		log.Fatal(err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := taskservice.NewTaskService(taskRepo)
	ctl := controller.NewTaskController(taskService)

	http.HandleFunc("/task", ctl.CreateTaskHandler)
	http.HandleFunc("/tasks", ctl.FindAllTaskHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
