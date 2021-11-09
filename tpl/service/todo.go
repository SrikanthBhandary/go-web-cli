package service

var Tmpl = `
package service

import (
	"{{.}}/entity"
	"{{.}}/repository"
)

type ToDoServiceInterface interface {
	FindAll() ([]entity.ToDo, error)
	FindByID(id int) (entity.ToDo, error)
	Delete() error
	Update() error
}

type ToDo struct {
	Repo repository.ToDoInterface
}

func NewToDoService(repo repository.ToDoInterface) ToDoServiceInterface {
	return &ToDo{Repo: repo}
}

func (service *ToDo) FindAll() ([]entity.ToDo, error) {
	return service.Repo.FindAll()
}

func (service *ToDo) FindByID(id int) (entity.ToDo, error) {
	return service.Repo.FindByID(id)
}

func (service *ToDo) Delete() error {
	return service.Repo.Delete()
}

func (service *ToDo) Update() error {
	return service.Repo.Update()
}
`
