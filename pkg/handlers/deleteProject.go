package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *BaseHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err.Error())
	}
	stmt, err := h.db.Prepare("DELETE FROM projects WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(200)
	displayMessage := "Deleted project id :" + params["id"]
	_, err = w.Write([]byte(displayMessage))
	if err != nil {
		panic(err.Error())
	}
}
