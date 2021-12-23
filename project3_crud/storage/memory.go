//go mod init
//go mod tidy

package storage

import (
	"github.com/andymartinezot/api_go_project/project3_crud/model"
)

//Memory estructura de memoria
type Memory struct {
	currentId int
	Persons   map[int]model.Person
}

//NewMemory devuelve una instancia de memory.
func NewMemory() Memory {
	//Inicialicia el mapa de persona
	persons := make(map[int]model.Person)

	return Memory{
		currentId: 0,
		Persons:   persons,
	}
}

func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentId++
	m.Persons[m.currentId] = *person

	return nil
}
