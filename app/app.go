package app

import (
	"fmt"
	"github.com/daftkid/mysportapp/domain/avatar"
	"github.com/daftkid/mysportapp/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	avatarHandler := AvatarHandlers{service.NewDefaultAvatarService(avatar.NewAvatarRepositoryDB())}

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/avatars", avatarHandler.getAllAvatars).Methods(http.MethodGet)
	router.HandleFunc("/avatars/{avatar_id:[0-9]+}", avatarHandler.getAvatarById).Methods(http.MethodGet)

	fmt.Println("starting")

	// starting a server
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
