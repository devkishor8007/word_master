package routes

import (
	"github.com/gorilla/mux"
	"github.com/devkishor8007/word_master/src/controllers"
	"github.com/devkishor8007/word_master/src/middleware"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/signin", controllers.Login).Methods("POST")
	r.Handle("/protected", middleware.RequiredAuth(http.HandlerFunc(controllers.Home))).Methods("GET")
	r.HandleFunc("/", controllers.Home).Methods("GET")
}