package core

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

const (
	cookieApp = "example-saludos-cookie"
)

var hashKey = []byte("1234567890123456")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("1234567890123456")
var s = securecookie.New(hashKey, blockKey)

type CookieRaw struct {
	Id    string
	User  string
	Email string
}

func NewCookie(raw CookieRaw) http.Cookie {
	var cookie http.Cookie
	if encoded, err := s.Encode(cookieApp, raw); err == nil {
		cookie = http.Cookie{
			Name:       cookieApp,
			Value:      encoded,
			Path:       "/",
			Domain:     "localhost",
			Expires:    time.Time{},
			RawExpires: "",
			MaxAge:     0,
			Secure:     false,
			HttpOnly:   false,
			SameSite:   0,
			Raw:        "",
			Unparsed:   nil,
		}
	}
	return cookie
}

func ExtractUserFromCookies(r *http.Request) (user CookieRaw) {
	if cookie, err := r.Cookie(cookieApp); err == nil {
		err = s.Decode(cookieApp, cookie.Value, &user)
	}
	return
}
