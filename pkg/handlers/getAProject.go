package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashu171996/Crud-rest-api/pkg/model"
	"github.com/gorilla/mux"
)

func (h *BaseHandler) GetAProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	row := h.db.QueryRow("SELECT id, name, country_code, manager_name from projects where id = ?", id)
	switch err := row.Scan(&project.ID, &project.Name, &project.CountryCode, &project.ManagerName); err {
	case sql.ErrNoRows:
		w.WriteHeader(200)
		displayMessage := "No rows were returned for requested id :" + params["id"]
		_, err := w.Write([]byte(displayMessage))
		if err != nil {
			panic(err.Error())
		}
	case nil:
		err = json.NewEncoder(w).Encode(project)
		if err != nil {
			panic(err.Error())
		}
	default:
		panic(err.Error())
	}
}
