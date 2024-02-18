package salutes

import (
	"github.com/gorilla/mux"
	"salu2/src/core"
)

func CreateUserSubrouter(v1 *mux.Router) {
	saluteSubrouter := core.Subrouter(v1, "/salutes")
	saluteSubrouter.HandleFunc("/get", GetSalutesByUser)
}
