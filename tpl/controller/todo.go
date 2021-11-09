package controller

var Tmpl = `
package controller

import (
	"net/http"	
	"encoding/json"

 	"{{.}}/service"	
)

type ToDo struct {	
	Svc    service.ToDoServiceInterface
	Config interface{}
}

func (server *ToDo) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (server *ToDo) FindAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := server.Svc.FindAll()
	json.NewEncoder(w).Encode(data)
}
`
