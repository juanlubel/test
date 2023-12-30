package salute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"salu2/backend/src/core"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	user := core.ExtractUserFromCookies(r)
	userRaw, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(userRaw); err != nil {
		return
	}
}

func LoginHandler(w http.ResponseWriter, _ *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("test", "hola")
	cookie := core.NewCookie(core.CookieRaw{
		User: "Juanlu",
	})

	http.SetCookie(w, &cookie)

	//cookie := http.Cookie{
	//	Name:     "exampleCookie",
	//	Value:    "Hello world!",
	//	Path:     "/",
	//	MaxAge:   3600,
	//	HttpOnly: true,
	//	Secure:   false,
	//}
	//
	//fmt.Println(cookie.String())
	//// Use the http.SetCookie() function to send the cookie to the client.
	//// Behind the scenes this adds a `Set-Cookie` header to the response
	//// containing the necessary cookie data.
	//http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"set cookie 2": true}`))
	if err != nil {
		return
	}
	//_, err := io.WriteString(w, `{"set cookie": true}`)
	//if err != nil {
	//	return
	//}
}
