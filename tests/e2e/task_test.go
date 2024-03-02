package e2e

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/application/services/taskservice"
	"TodoApp_basic/controller"
	"TodoApp_basic/domain/repository"
	"TodoApp_basic/testers"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Task_Create(b *testing.T) {

	//setup
	dns, db, err := testers.TestInit()
	if err != nil {
		b.Fatal(err)
	}
	defer testers.TestCleanUp(dns, db)

	//arrange
	mTask := model.NewTask("item1", false)
	taskRepo := repository.NewTaskRepository(db)
	taskService := taskservice.NewTaskService(taskRepo)

	// Define test cases
	benchmark := []struct {
		name         string
		requestBody  model.Task
		statusCode   int
		responseBody string //interface{}
		serviceErr   error
	}{
		{
			name:        "Valid taskusecase",
			requestBody: *mTask, //model.Task{Item: "valid item", Done: false},
			statusCode:  http.StatusCreated,
			//responseBody:  //map[string]interface{}{"message": "Task created successfully"}, // Replace with expected response
			serviceErr: nil,
		},
		{
			name:         "Empty item",
			requestBody:  model.Task{},
			statusCode:   http.StatusBadRequest,
			responseBody: "",
			serviceErr:   nil,
		},
		{
			name:         "Invalid Json",
			requestBody:  model.Task{Done: false},
			statusCode:   http.StatusBadRequest,
			responseBody: "",
			serviceErr:   nil,
		},
		{
			name:         "Service error",
			requestBody:  model.Task{Item: "valid item", Done: false},
			statusCode:   http.StatusInternalServerError,
			responseBody: "",
			serviceErr:   errors.New("some error"),
		},
	}

	for _, tc := range benchmark {
		b.Run(tc.name, func(t *testing.T) {
			// Create request
			reqBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			// Record the response
			w := httptest.NewRecorder()
			ctl := controller.NewTaskController(taskService)
			ctl.CreateTaskHandler(w, req)

			// Assert response status code
			assert.Equal(t, w.Code, tc.statusCode)

			// Assert response body (optional, adjust based on response format)
			if tc.statusCode == http.StatusCreated {
				var gotBody string //map[string]interface{}
				err = json.NewDecoder(w.Body).Decode(&gotBody)

				assert.NoError(t, err)
				assert.NotEmpty(t, gotBody)

			}
		})
	}

}
