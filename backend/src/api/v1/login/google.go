package login

import (
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"
	"net/http"
	"salu2/backend/src/core"
)

const (
	googleCookieApp = "example-saludos-cookie"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

var config = Config{
	ClientID:     "133549738783-he0u6rfko49uv556kpf2v7s4va74bdok.apps.googleusercontent.com",
	ClientSecret: "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",
}
var stateConfig = gologin.CookieConfig{
	Name:     "",
	Domain:   "",
	Path:     "",
	MaxAge:   0,
	HTTPOnly: false,
	Secure:   false,
	SameSite: 0,
}
var oauth2Config = &oauth2.Config{
	ClientID:     config.ClientID,
	ClientSecret: config.ClientSecret,
	RedirectURL:  "http://localhost:8080/v1/login/google/callback",
	Endpoint:     googleOAuth2.Endpoint,
	Scopes:       []string{"profile", "email"},
}

func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		googleUser, err := google.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 2. Implement a success handler to issue some form of session
		//session := sessionStore.New(sessionName)
		//session.Set(sessionUserKey, googleUser.Id)
		//session.Set(sessionUsername, googleUser.Name)
		//fmt.Println(googleUser.Id)
		//fmt.Println(googleUser.Name)

		cookie := core.NewCookie(core.CookieRaw{
			Id:    googleUser.Id,
			User:  googleUser.Name,
			Email: googleUser.Email,
		})

		http.SetCookie(w, &cookie)
		http.Redirect(w, req, "/v1/salute", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func GoogleSocialHandler() http.Handler {
	return google.StateHandler(stateConfig, google.LoginHandler(oauth2Config, nil))
}

func GoogleCallBackHandler() http.Handler {
	return google.StateHandler(stateConfig, google.CallbackHandler(oauth2Config, issueSession(), nil))
}
