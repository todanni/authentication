package repository

import (
	"github.com/todanni/template-repository/pkg/example"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) example.Repository  {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func (r repository) ExampleGet() ([]example.Example, error) {
	panic("implement me")
}
