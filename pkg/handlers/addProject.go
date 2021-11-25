package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ashu171996/Crud-rest-api/pkg/model"
)

func (h *BaseHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project
	stmt, err := h.db.Prepare("INSERT INTO projects(name, country_code, manager_name) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	json.NewDecoder(r.Body).Decode(&project)

	_, err = stmt.Exec(project.Name, project.CountryCode, project.ManagerName)
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New post was created")
}
