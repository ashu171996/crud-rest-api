package main

import (
	"log"
	"net/http"

	"github.com/ashu171996/Crud-rest-api/driver"
	"github.com/ashu171996/Crud-rest-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := driver.ConnectDB()
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	h := handlers.NewBaseHandler(db)
	router := mux.NewRouter().StrictSlash(true)

	// CRUD routing
	router.HandleFunc("/", h.ProjectHomePage).Methods("GET")
	router.HandleFunc("/project", h.GetProjects).Methods("GET")
	router.HandleFunc("/project", h.CreateProject).Methods("POST")
	router.HandleFunc("/project/{id}", h.GetAProject).Methods("GET")
	router.HandleFunc("/project/{id}", h.UpdateProject).Methods("PUT")
	router.HandleFunc("/project/{id}", h.DeleteProject).Methods("DELETE")

	log.Println("API is running! at http://localhost:4000")
	err = http.ListenAndServe(":4000", router)
	if err != nil {
		panic(err.Error())
	}
}
