package mockers

import "github.com/stretchr/testify/mock"

type MockId struct {
	mock.Mock
}

func (m *MockId) GetId() string {
	return m.Called().Get(0).(string)
}
