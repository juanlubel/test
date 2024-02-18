package salutes

import (
	"net/http"
)

func GetSalutesByUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("hola"))
	if err != nil {
		return
	}
}
