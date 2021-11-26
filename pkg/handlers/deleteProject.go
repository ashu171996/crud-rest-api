package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashu171996/crud-rest-api/pkg/model"
	"github.com/gorilla/mux"
)

// DeleteProject: this handle func used to delete a project
func (h *BaseHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	row := h.db.QueryRow("SELECT * from projects where id = ?", id)
	switch err := row.Scan(&project.ID, &project.Name, &project.CountryCode, &project.ManagerName); err {
	case sql.ErrNoRows:
		w.WriteHeader(200)
		displayMessage := model.Payload{
			Response: "No row was deleted for requested id " + params["id"],
		}
		err = json.NewEncoder(w).Encode(displayMessage)
		if err != nil {
			panic(err.Error())
		}
	case nil:
		stmt, err := h.db.Prepare("DELETE FROM projects WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(id)
		if err != nil {
			panic(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		displayMessage := model.Payload{
			Response: "Deleted project id :" + params["id"],
		}

		err = json.NewEncoder(w).Encode(displayMessage)
		if err != nil {
			panic(err.Error())
		}
		return
	default:
		panic(err.Error())
	}

}
