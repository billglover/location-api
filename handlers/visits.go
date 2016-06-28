package handlers

import (
	"github.com/gorilla/context"
	"gopkg.in/matryer/respond.v1"
	"net/http"
)

func VisitsGet(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotImplemented, nil)
	context.Clear(r)
}

func VisitsGetOne(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotImplemented, nil)
	context.Clear(r)
}

func VisitsPost(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusNotImplemented, nil)
	context.Clear(r)
}
