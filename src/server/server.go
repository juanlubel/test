package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"salu2/src/api/front"
	"salu2/src/api/v1/login"
	"salu2/src/api/v1/root"
	"salu2/src/api/v1/user"
	"salu2/src/core"
	"salu2/src/middleware"
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

func Start(config *Config) {
	r := Router()
	fmt.Printf("Server listening in port %v\n", config.Port)

	err := http.ListenAndServe(config.Port, r)
	if err != nil {
		panic(err)
	}
	//srv := &http.Server{
	//	Handler:      r,
	//	Addr:         "127.0.0.1" + config.Port,
	//	WriteTimeout: 15 * time.Second,
	//	ReadTimeout:  15 * time.Second,
	//}
	//log.Fatal(srv.ListenAndServe())

}

var dir = "./static/"

// Router sets up and configures the main router for the server.
// It creates a new server.Router, and then creates a versioned Subrouter using createVersionedSubrouter.
// Then it creates a salute Subrouter using createSaluteSubrouter.
// Finally, it returns the salute Subrouter, which is the main router for the server.
func Router() *mux.Router {
	fmt.Println("Preparing router")
	r := mux.NewRouter()

	front.CreateHomeSubrouter(r)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	v1 := createVersionedSubrouter(r, "/v1")
	v1.Use(middleware.CookieMiddleware)
	root.CreateSaluteSubrouter(v1)
	login.CreateLoginSubrouter(v1)
	user.CreateUserSubrouter(v1)
	return r
}

func createVersionedSubrouter(r *mux.Router, version string) *mux.Router {
	return core.Subrouter(r, version)
}
