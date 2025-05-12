package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/daftkid/mysportapp/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AvatarHandlers struct {
	service service.AvatarService
}

func (ah *AvatarHandlers) getAllAvatars(w http.ResponseWriter, r *http.Request) {
	avatars, _ := ah.service.GetAllAvatars()

	if r.Header.Get("Content-Type") == CONTENT_TYPE_XML {
		w.Header().Add("Content-Type", CONTENT_TYPE_XML)
		xml.NewEncoder(w).Encode(avatars)
	} else {
		w.Header().Add("Content-Type", CONTENT_TYPE_JSON)
		json.NewEncoder(w).Encode(avatars)
	}
}

func (ah *AvatarHandlers) getAvatarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["avatar_id"]
	avatar, err := ah.service.GetAvatarById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		if r.Header.Get("Content-Type") == CONTENT_TYPE_XML {
			w.Header().Add("Content-Type", CONTENT_TYPE_XML)
			xml.NewEncoder(w).Encode(avatar)
		} else {
			w.Header().Add("Content-Type", CONTENT_TYPE_JSON)
			json.NewEncoder(w).Encode(avatar)
		}
	}
}
