package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"salu2/backend/src/api/v1/login"
	"salu2/backend/src/api/v1/salute"
	"salu2/backend/src/core"
	"strings"
)

type Config struct {
	Port string
}

func NewConfig(port string) *Config {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return &Config{Port: port}
}

func New(config *Config) {
	r := Router()

	fmt.Printf("Server listening in port %v", config.Port)

	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		return
	}

}

// Router sets up and configures the main router for the server.
// It creates a new server.Router, and then creates a versioned Subrouter using createVersionedSubrouter.
// Then it creates a salute Subrouter using createSaluteSubrouter.
// Finally, it returns the salute Subrouter, which is the main router for the server.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hola", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hola"))
	})
	v1 := createVersionedSubrouter(r, "/v1")
	salute.CreateSaluteSubrouter(v1)
	login.CreateLoginSubrouter(v1)
	return v1
}

func createVersionedSubrouter(r *mux.Router, version string) *mux.Router {
	return core.Subrouter(r, version)
}
