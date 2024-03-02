package query

type IQuery interface {
	Add() string
	Update() string
	Delete() string
	FindOne() string
	FindAll() string
}

// Query type and properties
type Query struct {
	Add     string
	Update  string
	Delete  string
	FindOne string
	FindAll string
}
