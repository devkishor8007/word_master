package routes

import (
	"github.com/gorilla/mux"
	"github.com/devkishor8007/word_master/src/controllers"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/signin", controllers.Login).Methods("POST")
	r.Handle("/protected", controllers.RequiredAuth(http.HandlerFunc(controllers.Home))).Methods("GET")
	r.HandleFunc("/", controllers.Home).Methods("GET")
}