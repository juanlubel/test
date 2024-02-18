package front

import (
	"fmt"
	"net/http"
	"salu2/src/api/v1/user"
	"salu2/src/core"
	"salu2/src/pages/components"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	raw, err := core.ExtractUserFromCookies(r)
	if err != nil {
		fmt.Println(err)
	}
	newUser := user.NewUser(raw.Email, raw.Photo)
	newUser.Retrieve()
	w.WriteHeader(http.StatusOK)
	users := user.Users{}
	users.GetByGroup(newUser.MainGroup)
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
