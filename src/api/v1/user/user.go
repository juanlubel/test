package user

import (
	"fmt"
	"net/http"
	"salu2/src/core"
	"salu2/src/pages/components"
)

func GetAllUsersByGroup(w http.ResponseWriter, r *http.Request) {
	raw, err := core.ExtractUserFromCookies(r)
	if err != nil {
		fmt.Println(err)
	}
	user := NewUser(raw.Email, raw.Photo)
	user.Retrieve()
	w.WriteHeader(http.StatusOK)
	users := Users{}
	users.GetByGroup(user.MainGroup)
	var renderList []components.RenderUser
	for _, item := range users.List {
		render := components.RenderUser{
			Email:          item.Email,
			Avatar:         item.Avatar,
			LastConnection: item.LastConnection.String(),
		}
		renderList = append(renderList, render)
	}

	page := components.UserList(renderList)
	err = page.Render(r.Context(), w)
	if err != nil {
		return
	}
}

func PictureHandler(w http.ResponseWriter, r *http.Request) {
	user, err := core.ExtractUserFromCookies(r)
	if err != nil {
		fmt.Println(err)
	}
	avatar := components.Avatar(user.Photo)
	err = avatar.Render(r.Context(), w)
	if err != nil {
		return
	}
}

func WelcomeMessageHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	//w.Header().Set("Content-Type", "application/json")
	user, err := core.ExtractUserFromCookies(r)
	if err != nil {
		fmt.Println(err)
	}

	msg := GenerateRandomMessage(user.User)
	avatar := components.WelcomeMessage(msg)
	err = avatar.Render(r.Context(), w)
	if err != nil {
		return
	}
}

func GenerateRandomMessage(user string) string {
	message := "Bienvenido de nuevo, " + user
	return message
}
