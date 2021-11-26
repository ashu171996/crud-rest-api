package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ashu171996/crud-rest-api/pkg/model"
)

// GetProjects: this handle func used to get all projects
func (h *BaseHandler) GetProjects(w http.ResponseWriter, r *http.Request) {

	var projects []model.Project
	w.Header().Set("Content-Type", "application/json")

	result, err := h.db.Query("SELECT * from projects")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var project model.Project
		err := result.Scan(&project.ID, &project.Name, &project.CountryCode, &project.ManagerName)
		if err != nil {
			panic(err.Error())
		}
		projects = append(projects, project)
	}

	if len(projects) == 0 {
		displayMessage := model.Payload{
			Response: "There is no data available in projects",
		}
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(displayMessage)
		if err != nil {
			panic(err.Error())
		}
		return
	}
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		panic(err.Error())
	}
}
