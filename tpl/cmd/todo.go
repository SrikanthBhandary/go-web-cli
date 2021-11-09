package cmd

var Tmpl = `
package main

import (
	"os"
	stdlog "log"
	"net/http"

	admissioncontrol "github.com/elithrar/admission-control"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"{{.}}/controller"	
	"{{.}}/service"	
	"{{.}}/repository"	
)

var logger log.Logger

func main(){
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	logger = log.With(logger, "Level", "Info", "env", "DAILY", "ts", log.DefaultTimestampUTC)
	loggingMiddleware := admissioncontrol.LoggingMiddleware(logger)
	router := mux.NewRouter()
	loggedRouter := loggingMiddleware(router)
	repo := repository.NewToDoRepo(nil)
	server := controller.ToDo{service.NewToDoService(repo), nil}
	router.HandleFunc("/health", server.Health)
	router.HandleFunc("/todos", server.FindAll)
	logger.Log("Server started at port...", 8080)
	http.ListenAndServe(":8080", loggedRouter)

}
`
