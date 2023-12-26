package salute

import (
	"io"
	"net/http"
)

func RootHandler(w http.ResponseWriter, req *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
}
