package router

import (
	"TodoApp_basic/application/usecase/taskaction"
	"TodoApp_basic/controller"
	"TodoApp_basic/domain/repository"
	"TodoApp_basic/routes/logger"
	"TodoApp_basic/routes/usecase"
	"database/sql"
	"net/http"
	"time"
)

type Mux struct {
	log      logger.Logger
	port     Port
	storeSql *sql.DB
	ctx      time.Duration
	routes   []Route
	rep      repository.TaskRepository
	uc       usecase.Task
}

func NewMuxServer(logger logger.Logger, port Port, storeSql *sql.DB, ctxTimeout time.Duration) *Router {
	return &Router{
		log:      logger,
		port:     port,
		storeSql: storeSql,
		ctx:      ctxTimeout,
	}
}

var (
	mux = http.NewServeMux()

	rep  = repository.NewTaskRepository()
	uc   = taskaction.NewTaskService(rep)
	ctrl = controller.NewTaskController(uc)
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/")
}
