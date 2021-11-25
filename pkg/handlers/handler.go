package handlers

import "database/sql"

// BaseHandler will hold db connection
type BaseHandler struct {
	db *sql.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}
