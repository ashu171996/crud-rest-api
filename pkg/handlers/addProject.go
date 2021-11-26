package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ashu171996/crud-rest-api/pkg/model"
)

// CreateProject: this handle func used to create project
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
	displayMessage := model.Payload{
		Response: "Project created successfully",
	}

	err = json.NewEncoder(w).Encode(displayMessage)
	if err != nil {
		panic(err.Error())
	}
}
