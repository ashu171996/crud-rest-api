package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ashu171996/Crud-rest-api/pkg/model"
)

func (h *BaseHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var projects []model.Project
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
		w.Write([]byte("There is no data available in project"))
		w.WriteHeader(200)
		return
	}
	err = json.NewEncoder(w).Encode(projects)
	if err != nil {
		panic(err.Error())
	}
}
