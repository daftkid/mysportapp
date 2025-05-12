package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/daftkid/mysportapp/service"
	"net/http"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, _ := uh.service.GetAllUsers()

	if r.Header.Get("Content-Type") == CONTENT_TYPE_XML {
		w.Header().Add("Content-Type", CONTENT_TYPE_XML)
		xml.NewEncoder(w).Encode(users)
	} else {
		w.Header().Add("Content-Type", CONTENT_TYPE_JSON)
		json.NewEncoder(w).Encode(users)
	}
}
