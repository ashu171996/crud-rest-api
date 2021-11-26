package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashu171996/crud-rest-api/pkg/model"
)

// UpdateProject: this handle func used to update a project
func (h *BaseHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&project)
	stmt, err := h.db.Prepare("UPDATE projects SET name = ?, country_code = ?, manager_name = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(project.Name, project.CountryCode, project.ManagerName, project.ID)
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(200)
	id := strconv.Itoa(project.ID)
	displayMessage := model.Payload{
		Response: "Updated requested id :" + id,
	}
	err = json.NewEncoder(w).Encode(displayMessage)
	if err != nil {
		panic(err.Error())
	}
}
