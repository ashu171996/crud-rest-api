package main

import (
	"log"
	"net/http"

	"github.com/ashu171996/crud-rest-api/driver"
	"github.com/ashu171996/crud-rest-api/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// getting database connection
	db, err := driver.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	h := handlers.NewBaseHandler(db)
	router := mux.NewRouter().StrictSlash(true)

	// USing middleware
	router.Use(handlers.LoggingMiddleware)

	// GEt request
	router.HandleFunc("/project", h.GetProjects).Methods("GET")
	router.HandleFunc("/project/{id}", h.GetAProject).Methods("GET")

	// POST request
	router.HandleFunc("/project", h.CreateProject).Methods("POST")

	// PUT request
	router.HandleFunc("/project", h.UpdateProject).Methods("PUT")

	// DELETE request
	router.HandleFunc("/project/{id}", h.DeleteProject).Methods("DELETE")

	// Starting server
	log.Println("API is running! at http://localhost:4000")
	err = http.ListenAndServe(":4000", router)
	if err != nil {
		panic(err.Error())
	}
}
