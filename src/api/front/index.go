package front

import "github.com/gorilla/mux"

func CreateHomeSubrouter(front *mux.Router) {
	front.HandleFunc("/", HomeHandler)
	front.HandleFunc("/dashboard", ListHandler)
	front.HandleFunc("/users", UsersHandler)
	front.HandleFunc("/settings", SettingsHandler)
}
