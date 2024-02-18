package login

import (
	"fmt"
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"golang.org/x/oauth2"
	googleOAuth2 "golang.org/x/oauth2/google"
	"net/http"
	"salu2/src/api/v1/user"
	"salu2/src/core"
)

//var urlBase = os.Getenv("URL_BASE")

var urlBase = "http://localhost:8080"

//var urlBase = "https://saludos.spu-labs.dev"

type Config struct {
	ClientID     string
	ClientSecret string
}

var config = Config{
	ClientID:     "133549738783-he0u6rfko49uv556kpf2v7s4va74bdok.apps.googleusercontent.com",
	ClientSecret: "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",
}

var stateConfig = gologin.DebugOnlyCookieConfig

//	var stateConfig = gologin.CookieConfig{
//		Name:     "google_oauth",
//		Domain:   "",
//		Path:     "",
//		MaxAge:   0,
//		HTTPOnly: true,
//		Secure:   true,
//		SameSite: 0,
//	}
var oauth2Config = &oauth2.Config{
	ClientID:     config.ClientID,
	ClientSecret: config.ClientSecret,
	RedirectURL:  urlBase + "/v1/login/google/callback",
	//RedirectURL: "http://localhost:8080/v1/login/google/callback",
	Endpoint: googleOAuth2.Endpoint,
	Scopes:   []string{"profile", "email"},
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
			Photo: googleUser.Picture,
		})

		iuser := user.NewUser(googleUser.Email, googleUser.Picture)
		iuser.AddOne()

		http.SetCookie(w, &cookie)
		fmt.Println(googleUser)
		http.Redirect(w, req, "/", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func FailureHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/v1/error", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func GoogleSocialHandler() http.Handler {
	return google.StateHandler(stateConfig, google.LoginHandler(oauth2Config, nil))
}

func GoogleCallBackHandler() http.Handler {
	return google.StateHandler(stateConfig, google.CallbackHandler(oauth2Config, issueSession(), nil))
}
