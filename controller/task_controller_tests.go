package controller

/*
import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testers"

	"TodoApp_basic/application/model"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	gomock.mock
}

func (m *mockService) Add(taskusecase *model.Task) (interface{}, error) {
	args := m.Called(taskusecase)
	return args.Get(0), args.Error(1)
}

func TestCreateTaskHandler(t *testers.T) {
	// Setup mocks
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := &mockService{ctrl: ctrl}

	// Create handlers with mock service
	handlers := &TaskController{
		Service: mockService,
	}

	// Define test cases
	e2e := []struct {
		name         string
		requestBody  []byte
		statusCode   int
		responseBody interface{}
		serviceErr   error
	}{
		{
			name:         "Valid taskusecase",
			requestBody:  []byte(`{"Item": "Buy groceries"}`),
			statusCode:   http.StatusCreated,
			responseBody: map[string]interface{}{"message": "Task created successfully"}, // Replace with expected response
			serviceErr:   nil,
		},
		{
			name:         "Empty item",
			requestBody:  []byte(`{}`),
			statusCode:   http.StatusBadRequest,
			responseBody: nil,
			serviceErr:   nil,
		},
		{
			name:         "Invalid JSON",
			requestBody:  []byte(`{"invalid_json`),
			statusCode:   http.StatusInternalServerError,
			responseBody: nil,
			serviceErr:   nil,
		},
		{
			name:         "Service error",
			requestBody:  []byte(`{"Item": "Buy groceries"}`),
			statusCode:   http.StatusInternalServerError,
			responseBody: nil,
			serviceErr:   errors.New("some error"),
		},
	}

	for _, tc := range e2e {
		t.Run(tc.name, func(t *testers.T) {
			// Create request
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")

			// Set expectations on mock service
			mockService.EXPECT().Add(gomock.Any()).Return(tc.responseBody, tc.serviceErr)

			// Record the response
			w := httptest.NewRecorder()
			handlers.CreateTaskHandler(w, req)

			// Assert response status code
			if got, want := w.Code, tc.statusCode; got != want {
				t.Errorf("Unexpected status code: got %d, want %d", got, want)
			}

			// Assert response body (optional, adjust based on your response format)
			if tc.statusCode == http.StatusCreated {
				var gotBody map[string]interface{}
				err := json.NewDecoder(w.Body).Decode(&gotBody)
				if err != nil {
					t.Errorf("Error decoding response body: %v", err)
				}

				if !reflect.DeepEqual(gotBody, tc.responseBody) {
					t.Errorf("Unexpected response body: got %v, want %v", gotBody, tc.responseBody)
				}
			}
		})
	}
}
*/
