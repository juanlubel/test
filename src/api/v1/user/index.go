package user

import (
	"github.com/gorilla/mux"
	"salu2/src/core"
)

func CreateUserSubrouter(v1 *mux.Router) {
	userSubrouter := core.Subrouter(v1, "/user")
	userSubrouter.HandleFunc("/getAll", GetAllUsersByGroup)
	userSubrouter.HandleFunc("/picture", PictureHandler)
	userSubrouter.HandleFunc("/welcome", WelcomeMessageHandler)
}
