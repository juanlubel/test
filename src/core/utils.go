package core

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Subrouter(r *mux.Router, path string) *mux.Router {
	fmt.Println("\t" + path)
	return r.PathPrefix(path).Subrouter()
}
