package handlers

import (
	"net/http"
)

func (h *BaseHandler) ProjectHomePage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to CRUD API homepage"))
	if err != nil {
		panic(err.Error())
	}
}
