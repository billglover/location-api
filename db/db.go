package db

import (
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2"
	"net/http"
)

func WithDB(s *mgo.Session, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbcopy := s.Copy()
		defer dbcopy.Close()
		context.Set(r, "db", dbcopy)
		h(w, r)
	}
}
