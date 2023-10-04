package routes

import (
	"github.com/gorilla/mux"
	"github.com/devkishor8007/word_master/src/controllers"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/", controllers.Home).Methods("GET")
}