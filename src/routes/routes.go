package routes

import (
	"github.com/devkishor8007/word_master/src/controllers"
	"github.com/devkishor8007/word_master/src/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Register).Methods("POST")
	r.HandleFunc("/signin", controllers.Login).Methods("POST")
	r.Handle("/protected", middleware.RequiredAuth(http.HandlerFunc(controllers.Home))).Methods("GET")
	r.Handle("/profile", middleware.RequiredAuth(http.HandlerFunc(controllers.ViewProfile))).Methods("GET")
	r.HandleFunc("/", controllers.Home).Methods("GET")
}
