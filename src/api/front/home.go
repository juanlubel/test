package front

import (
	"net/http"
	"salu2/src/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.Home()
	err := page.Render(r.Context(), w)
	if err != nil {
		return
	}
}
