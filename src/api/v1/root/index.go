package root

import (
	"github.com/gorilla/mux"
)

// CreateSaluteSubrouter creates a subrouter for handling salute routes.
// It takes a mux.Router as an argument and returns nothing.
// The created subrouter is mounted at the path "/salute" on the provided router.
// It also registers the RootHandler as the handler function for the root path (empty path) of the subrouter.
//
// Example usage:
//
//	func Router() *mux.Router {
//		r := mux.NewRouter()
//		v1 := createVersionedSubrouter(r, "/v1")
//		salute.CreateSaluteSubrouter(v1)
//		return v1
//	}
func CreateSaluteSubrouter(v1 *mux.Router) {
	//saluteSubrouter := core.Subrouter(v1, "/")
	v1.HandleFunc("/", RootHandler)
	v1.HandleFunc("/error", ErrorHandler)
	v1.HandleFunc("/login", LoginHandler)
}
