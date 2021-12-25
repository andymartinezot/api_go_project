//go mod init
//go mod tidy

package storage

import (
	"fmt"

	"github.com/andymartinezot/api_go_project/project3_crud/model"
	//"$GOPATH/src/github.com/andymartinezot/api_go_project/project3_crud/model"
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

//Create crea una persona en el slice de memoria
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentId++
	m.Persons[m.currentId] = *person

	return nil
}

//Update actualiza una persona en el slice de memoria
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCanNotBeNil
	}

	//This "if" checks if the person exists through the ID in the map "Persons"
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesnotExists)
	}

	m.Persons[ID] = *person

	return nil
}

//Delete elimina a una persona en el slice de memoria
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesnotExists)
	}

	delete(m.Persons, ID)

	return nil
}

//GetByID retorna a una persona en el slice de memoria por el ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesnotExists)
	}

	return person, nil
}

//GetAll retorna todas las personas en el slice de memoria
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
