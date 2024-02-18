package front

import (
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	//user, err := core.ExtractUserFromCookies(r)
	//if user.User == "" {
	//	http.Redirect(w, r, "/v1/login/google", http.StatusMovedPermanently)
	//	return
	//}
	//page := components.SaluteList()
	//err := page.Render(r.Context(), w)
	//if err != nil {
	//	return
	//}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("dashboard"))
	if err != nil {
		return
	}
}
