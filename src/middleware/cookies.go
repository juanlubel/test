package middleware

import (
	"fmt"
	"net/http"
	"salu2/src/core"
	"strings"
)

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/v1/login") {
			next.ServeHTTP(w, r)
			return
		}
		// Do stuff here
		cookies, err := core.ExtractUserFromCookies(r)
		if cookies.User == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
