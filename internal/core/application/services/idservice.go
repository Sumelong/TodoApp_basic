package services

import (
	"github.com/google/uuid"
	"reflect"
)

type IdService struct {
	id string
}

func NewId() *IdService {
	id := uuid.NewString()
	return &IdService{id: id}
}

// SetId generates a new UUID string
func (s *IdService) SetId() {
	s.id = uuid.NewString()
}

func (s *IdService) GetId() string {
	return s.id
}

// GetVal function to get the value from the entity or model filed
func (s *IdService) GetVal(entity any, fieldName string) string {
	// Implement logic to extract the ID based on the entity type, e.g., using reflection
	return reflect.ValueOf(entity).FieldByName(fieldName).String()
}
