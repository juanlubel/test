package login

import (
	"github.com/gorilla/mux"
	"salu2/backend/src/core"
)

func CreateLoginSubrouter(v1 *mux.Router) {
	loginSubrouter := core.Subrouter(v1, "/login")
	loginSubrouter.Handle("/google", GoogleSocialHandler())
	loginSubrouter.Handle("/google/callback", GoogleCallBackHandler())
}
