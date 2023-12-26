package core

import "github.com/gorilla/mux"

func Subrouter(r *mux.Router, path string) *mux.Router {
	return r.PathPrefix(path).Subrouter()
}
