package controller

import (
	"TodoApp_basic/application/model"
	"TodoApp_basic/tests"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskController_CreateTaskHandler(t *testing.T) {

	//setup

	//arrange
	mockRepo := new(tests.MockRepository)
	mockService := NewMockTaskService(mockRepo)
	ctrl := NewTaskController(mockService)

	mTask := model.NewTask("item", false)
	reqBody, err := json.Marshal(mTask)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/task", bytes.NewReader(reqBody))
	if err != nil {
		t.Error(err)
	}
	rw := httptest.NewRecorder()

	//act
	ctrl.CreateTaskHandler(rw, req)

	//assert
	assert.HTTPSuccess(t, ctrl.CreateTaskHandler, "POST", "/task", nil)

}
