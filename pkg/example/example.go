package example

import (
	"gorm.io/gorm"
	"net/http"
)

type Example struct {
	ExampleString string `json:"example_string"`
	gorm.Model
}

// ExampleService - service definition with all methods attached to it
type Service interface {
	// ExampleMethod description
	ExampleMethod(w http.ResponseWriter, r *http.Request)
}

// ExampleRepository - repository definition with all the methods attached to retrieving data
type Repository interface {
	//ExampleGet returns a DB record of an example
	ExampleGet() ([]Example, error)
}