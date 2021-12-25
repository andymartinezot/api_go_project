package handler

import (
	"github.com/andymartinezot/api_go_project/project3_crud/model"
)

//Storage for the handler
type Storage interface {
	Create(person model.Person) error
	Update(ID int, person model.Person) error
	Delete(ID int) error
	GetById(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
