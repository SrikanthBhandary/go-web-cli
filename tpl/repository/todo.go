package repository

var Tmpl = `
package repository

import "{{.}}/entity"

type ToDoInterface interface {
	FindAll() ([]entity.ToDo, error)
	FindByID(id int) (entity.ToDo, error)
	Delete() error
	Update() error
}

type ToDo struct {
	DB interface{} //Actual DB Interface
}

func NewToDoRepo(db interface{}) ToDoInterface {
	return &ToDo{DB: db}
}

func (repo *ToDo) FindAll() ([]entity.ToDo, error) {
	todos := []entity.ToDo{
		{Id: 1, Title: "ToDo -1", Description: "Descrption -1"},
		{Id: 2, Title: "ToDo -2", Description: "Descrption -2"},
		{Id: 3, Title: "ToDo -3", Description: "Descrption -3"},
	}
	return todos, nil
}

func (repo *ToDo) FindByID(id int) (entity.ToDo, error) {
	return entity.ToDo{Id: 1, Title: "ToDo -1", Description: "Descrption -1"}, nil
}

func (repo *ToDo) Delete() error {
	return nil
}

func (repo *ToDo) Update() error {
	return nil
}
`
