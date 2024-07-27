package handlers

import (
	"quote/data"
	"net/http"
)

type handler struct {
	db data.Database
}

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Welcome(w http.ResponseWriter, r *http.Request)
}

func NewHandler(db data.Database) *handler {
	return &handler{db}
}
